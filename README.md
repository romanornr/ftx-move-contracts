# ftx-move-contracts

### MOVE Contracts
MOVE contracts are a straddle where the strike price is determined at the first hour of the day and expires at the last house of the day. (UTC).
So why did FTX decide to create MOVE contracts instead of only options like Deribit and Okcoin/Okex?
Options aren’t that popular under most retail traders, most retail traders in crypto have never tried trading Options, and FTX wants to provide centralized liquidity.

The downside at having so many different options with so many different strike prices and expirations on different exchanges is hard to market make for market makers and the liquidity, positions, and the risk is fractured. It gives a less liquid trading experience. FTX MOVE contracts are an attempt to create more collective knowledge of the contract to trade volatility.

Read my medium article about FTX MOVE Contracts
https://medium.com/@romanornr/ftx-com-move-contracts-46c586a66408

### What does this tool do?

This tool is for FTX daily MOVE contracts. It analyzes all daily MOVE Contracts from the current year we are in and shows the average expiration price of all daily MOVE Contracts. Not only that per also by weekday average and by months. For example, some days are more volatile than others. Data suggest that weekend days usually expire below the average expiration price. This tool can basically give you an edge.

Signup on FTX with this referral link receive a 10% discount instead of the regular 5% discount: https://ftx.com/#a=10percentDiscountOnFees


![alt text](https://github.com/romanornr/ftx-move-contracts/blob/master/screenshots/1.png?raw=true)
<br></br>

### Build from source (all platforms)

<details><summary><b>Install Dependencies</b></summary>

- **Go 1.13 or 1.14**

  Installation instructions can be found here: https://golang.org/doc/install.
  Ensure Go was installed properly and is a supported version:
  ```sh
  $ go version
  $ go env GOROOT GOPATH
  ```
  NOTE: `GOROOT` and `GOPATH` must not be on the same path. Since Go 1.8 (2016),
  `GOROOT` and `GOPATH` are set automatically, and you do not need to change
  them. However, you still need to add `$GOPATH/bin` to your `PATH` in order to
  run binaries installed by `go get` and `go install` (On Windows, this happens
  automatically).

  Unix example -- add these lines to .profile:

  ```
  PATH="$PATH:/usr/local/go/bin"  # main Go binaries ($GOROOT/bin)
  PATH="$PATH:$HOME/go/bin"       # installed Go projects ($GOPATH/bin)
  ```
  
  - **Git**

  Installation instructions can be found at https://git-scm.com or
  https://gitforwindows.org.
  ```sh
  $ git version
  ```
