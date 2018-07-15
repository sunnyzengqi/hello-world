package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

func main() {
	productPrefix := "insert into product (product_id,market_price,cost_price) values "
	productSuffix := " on duplicate key update market_price=values(market_price),cost_price=values(cost_price);"
	cashPrefix := "insert into product_ext (pid,cash) values "
	cashSuffix := " on duplicate key update cash=values(cash);"
	cashValues := ""
	productValues := ""
	excelFileName := "/Users/didi/Desktop/product-final.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if(i>1){
				productValues = productValues+ "("+row.Cells[0].String()+","+row.Cells[3].String()+","+row.Cells[4].String()+"),"
				cashValues = cashValues+ "("+row.Cells[0].String()+","+row.Cells[5].String()+"),"
			}
		}
	}
	productFinalValues := productValues[0:len(productValues)-1]
	cashFinalValues := cashValues[0:len(cashValues)-1]
	productUpdateSql := productPrefix+productFinalValues+productSuffix
	cashUpdateSql := cashPrefix+cashFinalValues+cashSuffix
	fmt.Println("product updatesql: "+productUpdateSql)
	fmt.Println("cash updatesql: "+cashUpdateSql)
	changeStatus()
}

func changeStatus(){
	excelFileName := "/Users/didi/Desktop/changeStatus.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			rowInt,_ := row.Cells[2].Int64()
			sql := "update pay_order_"+strconv.FormatInt(rowInt%16,10)+" set order_status=8 where order_id = '"+ row.Cells[3].String()+"';"
			fmt.Println(sql)
			}
			fmt.Println("len: ",len(sheet.Rows))
	}
}