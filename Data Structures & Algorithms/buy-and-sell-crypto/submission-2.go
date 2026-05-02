func maxProfit(prices []int) int {
    minPrice:=prices[0]
    maxProfite:=0
    for _, value:= range prices{

        if minPrice>value{
            minPrice=value
        }

        if value-minPrice> maxProfite {
            maxProfite=value-minPrice

        }
    }

    return maxProfite
}
