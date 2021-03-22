# smscrubbing
SMS Template scrubbing - India DLT

Golang , DLT SMS/Message content template validation library.

## DLT Content Template Validations are as below,

* Each variable length can’t exceed 30 chars
* Variable format is {#var#} which is case sensitive
* Promotional SMS content without consent template wont get approved.
* Variable can be inserted by clicking the radio button (insert variable)
* Two or more spaces are not supposed to use between 2 words, before word or after word.
* All special characters (found on keyboard) are allowed, except < and > symbols.
* Trans/Service category messages should have variable mandatorily.
* Promo category can have complete fixed content or with variable part
* Values like amount, date, a/c no, merchant names, OTP, codes, URL, customer names, cardtype, etc. needs to be replaced with variables.

Validation is done using above rules. 
