package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*

DLT Content Template Validations are as below,

Each variable length can’t exceed 30 chars
Variable format is {#var#} which is case sensitive
Promotional SMS content without consent template wont get approved.
Variable can be inserted by clicking the radio button (insert variable)
Two or more spaces are not supposed to use between 2 words, before word or after word.
All special characters (found on keyboard) are allowed, except < and > symbols.
Trans/Service category messages should have variable mandatorily.
Promo category can have complete fixed content or with variable part
Values like amount, date, a/c no, merchant names, OTP, codes, URL, customer names, cardtype, etc. needs to be replaced with variables.

*/

func main() {
	var inputdata [18]string
	inputdata[0] = "824926 is the otp for trxn of inr 57.75 at zaak epayment services pv with your sbi card xx3931. otp is valid for 10 mins. pls do not share with anyone"         // Assign a value to the first element
	inputdata[1] = "032456 is your OTP for fund transfer for amount Rs.3,000 to Ravi. OTP valid for 8 minutes. Do not share this OTP with anyone."                                  // Assign a value to the second element
	inputdata[2] = "428684 is OTP for your eComm Txn for amount Rs.15,000 OTP valid for 8 minutes. Do not share this OTP with anyone."                                              // Assign a value to the third element
	inputdata[3] = "369147 is OTP for your premium payment for amount Rs.34,000. OTP valid for 8 minutes. Do not share this OTP with anyone."                                       // Assign a value to the third element
	inputdata[4] = "852456 is your OTP for BillDesk Payment in NetBanking. OTP is valid for 8 minutes."                                                                             // Assign a value to the third element
	inputdata[5] = "Thank you for using EMI Facility on your IDBI Bank Credit Card 4***3495 EMI request for Rs. 25000.00 executed on 01/07/2019"                                    // Assign a value to the third element
	inputdata[6] = "YES BANK - Your new bill for BESCOM Bangalore - account 0842948000 for Rs 4339.00 could not get scheduled because auto pay limit is less than the bill amount." // Assign a value to the third element
	inputdata[7] = "account: 674508 is your samsung account verification code."
	inputdata[8] = "transaction alert: 49.0 was used from your flipkart gift card 6000172013334850 for order od117666705985700000 on flipkart.balance remaining in the card: 0.0.if you dont recognise this transaction, please reach out to http://fkrt.it/q0rbconnnn immediately."
	inputdata[9] = "Kindly note that the free look period for your insurance cancellation is 15 days from date of receipt of insurance policy. Regards, Bajaj Finance Ltd."
	inputdata[10] = "dear k, otp is 2568 for order id #101794788 at dailyorders phone case maker mobile app, kindly enter it to confirm your order. thank you!"
	inputdata[11] = "Your Rs.500 exclusive voucher is UNUSED!! Redeem it on purchase of Rs.1,000 at Marks & Spencer. Use code 654321001 Valid till 31st Mar 2020! T&C."
	inputdata[12] = "Bajaj Finserv Personal Loan needs Minimal Documentation. Fulfil your financial needs in one click http://m.BajFin.in/Iphr8tFE."
	inputdata[13] = "swedish massage 60min(1): 1355.93,aroma massage 60min(1): 1525.42,s.total(2): 2881.35,"
	inputdata[14] = "Hi, In order to best serve you and others, could you click on mosl.co/ywq8FBJpAn to share your meeting experience with Motilal Oswal RM Raju Saha on 22"
	inputdata[15] = `Lifetime Free ICICI Bank Credit Card with Vouchers from LensKart, Pepperfry, Grabon worth Rs.3000. SMS "apply" to 5676766 TnC apply`
	inputdata[16] = "Pay JUST Rs 640* pm & get Rs 83,333 for 120 months or payout of Rs 1,00,00,000 With LIC*(Life Insurance Cover) For Your Family. http://px2.in/pAD4Tls"
	inputdata[17] = "YOU can win Rs 20,000 in Fantasy cricket use code 542321. Install Qureka Pro app now to WIN Click - https://abc.com"

	var templates [18]string
	templates[0] = "{#var#} is the otp for trxn of inr {#var#} at {#var#} with your sbi card{#var#}. otp is valid for {#var#}. pls do not share with anyone"
	templates[1] = "{#var#} is your OTP for fund transfer for amount {#var#} to {#var#}. OTP valid for 8 minutes. Do not share this OTP with anyone."
	templates[2] = "{#var#} is OTP for your eComm Txn for amount {#var#} OTP valid for 8 minutes. Do not share this OTP with anyone."
	templates[3] = "{#var#} is OTP for your premium payment for amount {#var#}. OTP valid for 8 minutes. Do not share this OTP with anyone."
	templates[4] = "{#var#} is your OTP for BillDesk Payment in NetBanking. OTP is valid for 8 minutes."
	templates[5] = "Thank you for using EMI Facility on your IDBI Bank Credit Card {#var#} EMI request for {#var#} executed on {#var#}"
	templates[6] = "YES BANK - Your new bill for {#var#} - account {#var#} for Rs {#var#} could not get scheduled because auto pay limit is less than the bill amount."
	templates[7] = "account: {#var#} is your samsung account verification code."
	templates[8] = "transaction alert: {#var#} was used from your {#var#} gift card {#var#} for order {#var#} on flipkart.balance remaining in the card: {#var#}.if you dont recognise this transaction, please reach out to {#var#} immediately."
	templates[9] = "Kindly note that the free look period for your insurance cancellation is {#var#} from date of receipt of insurance policy. Regards, Bajaj Finance Ltd."
	templates[10] = `dear {#var#}, otp is {#var#} for order id {#var#} at dailyorders phone case maker mobile app, kindly enter it to confirm your order. thank you!`
	templates[11] = `Your Rs.{#var#} exclusive voucher is UNUSED!! Redeem it on purchase of Rs.{#var#} at Marks & Spencer. Use code {#var#} Valid till {#var#}! T&C.`
	templates[12] = "Bajaj Finserv Personal Loan needs Minimal Documentation. Fulfil your financial needs in one click {#var#}."
	templates[13] = "swedish massage {#var#},aroma massage {#var#},s.total{#var#},"
	templates[14] = "Hi, In order to best serve you and others, could you click on {#var#} to share your meeting experience with {#var#}"
	templates[15] = `Lifetime Free ICICI Bank Credit Card with Vouchers from LensKart, Pepperfry, Grabon worth Rs.{#var#}. SMS "{#var#}" to 5676766. TnC apply`
	templates[16] = "Pay JUST Rs {#var#} pm & get Rs {#var#} for {#var#} months or payout of Rs {#var#} With LIC (Life Insurance Cover) For Your Family. {#var#}"
	templates[17] = "YOU can win Rs {#var#} in Fantasy cricket use code {#var#}. Install Qureka Pro app now to WIN Click - {#var#}"

	for ii, input := range inputdata {
		match, _ := regexp.MatchString(`^.*[\s<>]{2,}.*$`, input)
		if match {
			fmt.Println("Input content has more than 2 spaces or < , > char.")
			continue
		}

		for ti, template := range templates {
			regex := strings.Replace("^"+template+"$", "{#var#}", ".{1,30}", -1)
			match, _ := regexp.MatchString(regex, input)
			if match {
				fmt.Println("Matched input index :", ii, "template index : ", ti)
				continue
			}
		}
	}
}
