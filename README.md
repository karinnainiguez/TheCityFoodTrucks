# TheCityFoodTrucks

Welcome to the Command Line Interface tool designed to find open food trucks in the city of San Francisco. 

## Usage

This tool is built using Golang.  You will need to have [Go installed on your computer, your Go Workspace created, and the path to that workspace stored in your ``GOPATH`` ](https://www.callicoder.com/golang-installation-setup-gopath-workspace/). 

When go has been completed, ensure that the following two variables are set up in your bash profile: 

```
export GOPATH=/usr/lib/go/

export PATH=$PATH:$GOPATH/bin
```

Clone this repository in your GOPATH.  Within the TheCityFoodTrucks folder, run the following commands to install dependencies and the tool: 

```
 🍕 /TheCityFoodTrucks $ go get -d ./...
 🍕 /TheCityFoodTrucks $ go build
 🍕 /TheCityFoodTrucks $ go install
```

After installation, you may use the ```$ TheCityFoodTrucks ``` command in order to see food trucks in San Francisco currently open.  The results are paginated in groups of 10 trucks per page, and the table will display the truck name, address, and times of operation.

Below are some examples of users running this command at different times: 


```
2018-11-19T10:29:48-08:00 [✿]  There are a total of 178 open food trucks in San Francisco now.
2018-11-19T10:29:48-08:00 [✿]      page 1 of 18: 

+--------------------------------+----------------------+---------------------+
|           TRUCK NAME           |       ADDRESS        |        TIMES        |
+--------------------------------+----------------------+---------------------+
| Anas Goodies Catering          | 1411 MARKET ST       | Monday  10AM - 11AM |
| Anas Goodies Catering          | 2150 CESAR CHAVEZ ST | Monday  10AM - 11AM |
| Anas Goodies Catering          | 368 ELM ST           | Monday  10AM - 11AM |
| Anas Goodies Catering          | 150 OTIS ST          | Monday  10AM - 11AM |
| Athena SF Gyro                 | 699 08TH ST          | Monday  6AM - 6PM   |
| BH & MT LLC                    | 2145 MARKET ST       | Monday  10AM - 11AM |
| BH & MT LLC                    | 3253 16TH ST         | Monday  10AM - 11AM |
| BH & MT LLC                    | 1477 GROVE ST        | Monday  10AM - 11AM |
| BOWL'D ACAI, LLC.              | 451 MONTGOMERY ST    | Monday  9AM - 3PM   |
| Bay Area Mobile Catering, Inc. | 1301 CESAR CHAVEZ ST | Monday  9AM - 5PM   |
| dba. Taqueria Angelica's       |                      |                     |
+--------------------------------+----------------------+---------------------+
Use the arrow keys to navigate: ↓ ↑ → ← 
? Page No. 1  Options: 
  ▸ Next
    Exit
```
Above, the user ran the command at 10:29am on a Monday.  Since there are 178 food trucks open, the user has the option to go navigate through the pages of food trucks by selecting "Next" from the options list.

```
2018-11-19T10:30:49-08:00 [✿]  There are a total of 178 open food trucks in San Francisco now.
2018-11-19T10:30:49-08:00 [✿]      page 3 of 18: 

+--------------------------+--------------------+---------------------+
|        TRUCK NAME        |      ADDRESS       |        TIMES        |
+--------------------------+--------------------+---------------------+
| D & T Catering           | 3100 PACIFIC AVE   | Monday  10AM - 11AM |
| D & T Catering           | 3839 WASHINGTON ST | Monday  10AM - 11AM |
| D & T Catering           | 2970 BROADWAY      | Monday  10AM - 11AM |
| D & T Catering           | 2840 BROADWAY      | Monday  10AM - 11AM |
| D & T Catering           | 2201 VALLEJO ST    | Monday  10AM - 11AM |
| DO UC US Mobile Catering | 400 PARNASSUS AVE  | Monday  9AM - 4PM   |
| El Gallo Jiro            | 3055 23RD ST       | Monday  10AM - 6PM  |
| El Tonayanse #4 / #36    | 1717 HARRISON ST   | Monday  10AM - 7PM  |
| El Tonayense #60         | 401 TREAT AVE      | Monday  10AM - 4PM  |
| Eli's Hot Dogs           | 101 BAY SHORE BLVD | Monday  9AM - 5PM   |
+--------------------------+--------------------+---------------------+
Use the arrow keys to navigate: ↓ ↑ → ← 
? Page No. 3  Options: 
  ▸ Previous
    Next
    Exit
```
As the user explores other pages of results, an option for "Previous" page will also appear.  At any point, the user may select "Exit" to exit the program.


**NOTE:**
When the command is run, and there are less than 10 food trucks open at the time, the user will not see "Previous" or "Next" options, as there is only one page of results.


## Implementation

This tool is built with [Golang](https://golang.org/).  The data regarding food trucks is provided by the [city of San Francisco](https://data.sfgov.org/Economy-and-Community/Mobile-Food-Schedule/jjew-r69b), with the help of the [community SDK for Socrata](https://github.com/SebastiaanKlippert/go-soda) when making get requests to that API. 


## Limitations

Right now, the user needs to be comfortable using the command line in order to use the tool.  Additionally, they will only see food trucks open at the time of running the program. 

## Possible Improvements

In order to address some of the limitations, it would be helpful to allow the user to set a flag with a date and time.  That would allow someone to plan ahead and see what food trucks are open in the future, or what food trucks were open on a specific date in the past.  We could also expand the usage of the tool by providing a web based user interface, to allow a larger audience to use the tool, including those users that are not comfortable using the command line.  By providing a user friendly web UI, we would also be helping users that ARE comfortable with the command line, by streamlining that usage.  Right now, the tool requires the user to have Goland installed, Go variables set, the project cloned or copied, and dependencies downloaded before they can even install it.  These steps may seem trivial if someone is already using Golang and has some of those steps done, but would mean there is a barrier to entry for using our tool, and the involved process might detour people from even using TheCityFoodTrucks themselves.  Building that web UI would require time and patience, as we would have to design, implement, and test an entire new portion of the program.  It would also involve business decisions about best color scheme, font, logo, and even how buttons are displayed.  Additionally, we'd have to make sure that the tool is accessible through a wide range of web browsers, and if it uses web browser plugins, that we tell the user so that they may have the best experience possible.

As always, before we begin making new features and enhancements, it would be beneficial to have full code coverage through testing.  That would ensure that we do not break current functionality when implementing anything else. 
