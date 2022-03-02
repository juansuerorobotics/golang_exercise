
Please follow the directions below to verfiy the code works.
below are 2 sections.  1st section has steps for running the test yourself.
2nd section is output from running the steps myself on my computer

==========================================================================



TEST RUN #1

STEP1: run the program with the command 

go run ibmcodechallenge.go 


STEP2:  cut and paste to the command line:

1 book at 12.49
1 music CD at 14.99
1 chocolate bar at 0.85

STEP3: press enter twice.

-------------------------------------------

TEST RUN #2

STEP1: run the program with the command 

go run ibmcodechallenge.go 


STEP2:  cut and paste to the command line:

1 imported box of chocolates at 10.00
1 imported bottle of perfume at 47.50


STEP3: press enter twice.

-------------------------------------------

TEST RUN #2

STEP1: run the program with the command 

go run ibmcodechallenge.go 


STEP2:  cut and paste to the command line:

1 imported bottle of perfume at 27.99
1 bottle of perfume at 18.99
1 packet of headache pills at 9.75
1 box of imported chocolates at 11.25


STEP3: press enter twice.






==========================================================================


THE FOLLOWING IS A SAMPLE RUN ON MY COMPUTER.


TEST RUN #1

metadojo@metadojoai:~/org/juansuero_ibm_codechallenge$ go run ibmcodechallenge.go 

IBM Code Challenge for dedellis@us.ibm.com.........

PROBLEM TWO: SALES TAXES

	usage:   numitems [imported] product name at price 
	example: 1 imported box of chocolates at 10.00 
	( press enter twice after the last item to print receipt ) 


1 book at 12.49
1 music CD at 14.99
1 chocolate bar at 0.85

1 book: 12.49
1 music CD: 16.49
1 chocolate bar: 0.85
Sales Taxes: 1.50
Total: 29.83




--------------------------

TEST RUN #2

metadojo@metadojoai:~/org/juansuero_ibm_codechallenge$ go run ibmcodechallenge.go 

IBM Code Challenge for dedellis@us.ibm.com.........

PROBLEM TWO: SALES TAXES

	usage:   numitems [imported] product name at price 
	example: 1 imported box of chocolates at 10.00 
	( press enter twice after the last item to print receipt ) 


1 imported box of chocolates at 10.00
1 imported bottle of perfume at 47.50

1 imported box of chocolates: 10.50
1 imported bottle of perfume: 54.65
Sales Taxes: 7.65
Total: 65.15



--------------------------

TEST RUN #3

metadojo@metadojoai:~/org/juansuero_ibm_codechallenge$ go run ibmcodechallenge.go 

IBM Code Challenge for dedellis@us.ibm.com.........

PROBLEM TWO: SALES TAXES

	usage:   numitems [imported] product name at price 
	example: 1 imported box of chocolates at 10.00 
	( press enter twice after the last item to print receipt ) 


1 imported bottle of perfume at 27.99
1 bottle of perfume at 18.99
1 packet of headache pills at 9.75
1 box of imported chocolates at 11.25

1 imported bottle of perfume: 32.19
1 bottle of perfume: 20.89
1 packet of headache pills: 9.75
1 imported box of  chocolates: 11.85
Sales Taxes: 6.70
Total: 74.68