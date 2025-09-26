package limitOrderTxTC

import (
	"log"
	"math/big"
	"math/rand"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/core/orderbook"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/myzhan/boomer"
)

const Name = "limitOrderTxTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	cursor uint32

	// User settings
	baseToken  = "2"
	quoteToken = "3"

	marketRules = orderbook.NewMarketRules()
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

	from := accGrp[atomic.AddUint32(&cursor, 1)%uint32(nAcc)]

	start := boomer.Now()
	err := SendRandomTx(cli, from)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.RecordSuccess("http", "SendNewLimitOrderTx"+" to "+endPoint, elapsed, int64(10))
	} else {
		boomer.RecordFailure("http", "SendNewLimitOrderTx"+" to "+endPoint, elapsed, err.Error())
	}
}

func SendRandomTx(cli *ethclient.Client, from *account.Account) error {
	var (
		side      orderbook.Side
		price     *big.Int
		quantity  *big.Int
		tx        *types.Transaction
		orderType = orderbook.LIMIT
		err       error
	)

	// Generate random price between 0.95 and 1.05 (with 18 decimals)
	// 0.95 * 1e18 = 950000000000000000
	// 0.05 * 1e18 = 50000000000000000 (each side range)
	// Total range: 0.1 * 1e18 = 100000000000000000
	minPrice := new(big.Int).Mul(big.NewInt(95), big.NewInt(1e16)) // 0.95 * 1e18
	priceRange := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e16)) // 0.1 * 1e18
	randomOffset := new(big.Int).Rand(rand.New(rand.NewSource(rand.Int63())), priceRange)
	price = new(big.Int).Add(minPrice, randomOffset)

	// Get tick size and adjust price to be compliant with tick
	priceUint256, _ := uint256.FromBig(price)
	tick := marketRules.GetTickSize(priceUint256)
	if tick != nil && tick.Sign() > 0 {
		// Round price to nearest tick
		tickBig := tick.ToBig()
		remainder := new(big.Int).Mod(price, tickBig)
		if remainder.Sign() > 0 {
			price.Sub(price, remainder)
			// If remainder > tick/2, round up
			halfTick := new(big.Int).Div(tickBig, big.NewInt(2))
			if remainder.Cmp(halfTick) > 0 {
				price.Add(price, tickBig)
			}
		}
	}

	// Generate random quantity between 2 and 3 (with 18 decimals)
	// 2 * 1e18 = 2000000000000000000
	// 1 * 1e18 = 1000000000000000000 (range)
	minQuantity := new(big.Int).Mul(big.NewInt(2), big.NewInt(1e18)) // 2 * 1e18
	quantityRange := new(big.Int).SetInt64(1e18) // 1 * 1e18
	randomOffset = new(big.Int).Rand(rand.New(rand.NewSource(rand.Int63())), quantityRange)
	quantity = new(big.Int).Add(minQuantity, randomOffset)

	// Get lot size and adjust quantity to be compliant with lot
	lot := marketRules.GetLotSize(priceUint256)
	if lot != nil && lot.Sign() > 0 {
		// Round quantity to nearest lot
		lotBig := lot.ToBig()
		remainder := new(big.Int).Mod(quantity, lotBig)
		if remainder.Sign() > 0 {
			quantity.Sub(quantity, remainder)
			// If remainder > lot/2, round up
			halfLot := new(big.Int).Div(lotBig, big.NewInt(2))
			if remainder.Cmp(halfLot) > 0 {
				quantity.Add(quantity, lotBig)
			}
		}
	}

	// 50/50 probability for BUY/SELL
	if rand.Intn(2) == 0 {
		side = orderbook.BUY
	} else {
		side = orderbook.SELL
	}

	tx, err = from.GenNewOrderTx(baseToken, quoteToken, side, price, quantity, orderType)
	if err != nil {
		log.Printf("Failed to generate new order tx: error=%v, baseToken=%s, quoteToken=%s, side=%d, price=%s, quantity=%s, orderType=%d",
			err, baseToken, quoteToken, side, price.String(), quantity.String(), orderType)
		return err
	}

	_, err = from.SendTx(cli, tx)
	if err != nil {
		log.Printf("Failed to send new order tx: error=%v, baseToken=%s, quoteToken=%s, side=%d, price=%s, quantity=%s, orderType=%d\n",
			err, baseToken, quoteToken, side, price.String(), quantity.String(), orderType)
	}
	return err
}

func scaleUp(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(1e18))
}
