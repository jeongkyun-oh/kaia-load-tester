package limitOrderNoTradeTxTC

import (
	"log"
	"math/big"
	"math/rand"
	"sync/atomic"

	obtypes "github.com/ethereum/go-ethereum/core/orderbook/v2/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/myzhan/boomer"
)

const Name = "limitOrderNoTradeTxTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	cursor uint32

	// User settings
	baseToken  = "3"
	quoteToken = "2"
	orderType  = obtypes.LIMIT
)

func Init(accs []*account.Account, endpoint string, _ *big.Int) {
	endPoint = endpoint

	cliCreate := func() any {
		c, err := ethclient.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	cliPool.Init(1000, 3000, cliCreate)

	accGrp = append(accGrp, accs...)

	nAcc = len(accGrp)
}

func Run() {
	cli := cliPool.Alloc().(*ethclient.Client)
	defer cliPool.Free(cli)

	var (
		from     = accGrp[atomic.AddUint32(&cursor, 1)%uint32(nAcc)]
		quantity = scaleUp(1)
		side     = obtypes.OrderSide(rand.Intn(2))
		price    = scaleUp(3 - int64(side)) // If buy, $3. If sell, $2.
	)

	start := boomer.Now()
	tx, err := from.GenNewOrderTx(baseToken, quoteToken, side, price, quantity, orderType)
	if err != nil {
		log.Printf("Failed to generate new order tx: error=%v, baseToken=%s, quoteToken=%s, side=%d, price=%s, quantity=%s, orderType=%d",
			err, baseToken, quoteToken, side, price.String(), quantity.String(), orderType)
		return
	}
	_, err = from.SendTx(cli, tx)
	elapsed := boomer.Now() - start
	if err != nil {
		log.Printf("Failed to send new order tx: error=%v, baseToken=%s, quoteToken=%s, side=%d, price=%s, quantity=%s, orderType=%d\n",
			err, baseToken, quoteToken, side, price.String(), quantity.String(), orderType)
		boomer.RecordFailure("http", "SendNewOrderTx"+" to "+endPoint, elapsed, err.Error())
		return
	}

	boomer.RecordSuccess("http", "SendNewOrderTx"+" to "+endPoint, elapsed, int64(10))
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}
