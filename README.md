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
2018-11-16T17:34:42-08:00 [✿]  There are a total of 43 open food trucks in San Francisco now.
2018-11-16T17:34:42-08:00 [✿]      page 1 of 5: 

+--------------------------+--------------------+--------------------+
|        TRUCK NAME        |      ADDRESS       |       TIMES        |
+--------------------------+--------------------+--------------------+
| Mob Dog                  | 720 MARKET ST      | Friday  8AM - 10PM |
| Tacos El Flaco           | 2901 03RD ST       | Friday  7AM - 6PM  |
| Yummy Hot Dogs           | 945 MARKET ST      | Friday  10AM - 7PM |
| Golden Gate Halal Food   | 1169 MARKET ST     | Friday  8AM - 8PM  |
| Halal Cart, LLC          | 901 MARKET ST      | Friday  10AM - 8PM |
| Singh Brothers Ice Cream | 360 PENINSULA AVE  | Friday  5PM - 6PM  |
| Alfaro Truck             | 306 VALENCIA ST    | Friday  2PM - 10PM |
| CC Acquisition LLC       | 298 MARKET ST      | Friday  7AM - 6PM  |
| Kettle Corn Star         | 865 MARKET ST      | Friday  10AM - 6PM |
| San Pancho's Tacos       | 491 BAY SHORE BLVD | Friday  6AM - 10PM |
+--------------------------+--------------------+--------------------+
Use the arrow keys to navigate: ↓ ↑ → ← 
? Page No. 1  Options: 
  ▸ Next
    Exit
```
Above, the user ran the command at 5:34pm on a Friday.  Since there are 43 food trucks, the user has the option to go navigate through the pages of food trucks by selecting "Next" from the options list.

```
2018-11-16T17:34:57-08:00 [✿]  There are a total of 43 open food trucks in San Francisco now.
2018-11-16T17:34:57-08:00 [✿]      page 3 of 5: 

+-------------------+----------------------+---------------------+
|    TRUCK NAME     |       ADDRESS        |        TIMES        |
+-------------------+----------------------+---------------------+
| Tacos Santana     | 2142 JERROLD AVE     | Friday  5PM - 10PM  |
| The Chai Cart     | 79 NEW MONTGOMERY ST | Friday  7AM - 6PM   |
| Linda's Catering  | 260 TOWNSEND ST      | Friday  12PM - 10PM |
| Got Snacks        | 1020 03RD ST         | Friday  10AM - 9PM  |
| Tacos Rodriguez   | 600 MENDELL ST       | Friday  8AM - 6PM   |
| La Jefa           | 531 BAY SHORE BLVD   | Friday  7AM - 8PM   |
| Linda's Catering  | 999 DIVISADERO ST    | Friday  7AM - 7PM   |
| Scotch Bonnet     | 115 SANSOME ST       | Friday  6AM - 8PM   |
| Giant Burrito     | 353 BAY SHORE BLVD   | Friday  7AM - 6PM   |
| Santana ESG, Inc. | 200 SHOTWELL ST      | Friday  10AM - 10PM |
+-------------------+----------------------+---------------------+
Use the arrow keys to navigate: ↓ ↑ → ← 
? Page No. 3  Options: 
  ▸ Previous
    Next
    Exit

```
As the user explores other pages of results,an option for "Previous" page will also appear.  At any point, the user may select "Exit" to exit the program.


**NOTE:**
When the command is run, and there are less than 10 food trucks open at the time, the user will not see "Previous" or "Next" options, as there is only one page of results.


## Implementation

This tool is built with [Golang](https://golang.org/).  The data regarding food trucks is provided by the [city of San Francisco](https://data.sfgov.org/Economy-and-Community/Mobile-Food-Schedule/jjew-r69b), with the help of the [community SDK for Socrata](https://github.com/SebastiaanKlippert/go-soda) when making get requests to that API. 


## Limitations

Right now, the user needs to be comfortable using the command line in order to use the tool.  Additionally, they will only see food trucks open at the time of running the program. 

## Possible Improvements

In order to address some of the limitations, it would be helpful to allow the user to set a flag with a date and time.  That would allow someone to plan ahead and see what food trucks are open in the future, or what food trucks were open on a specific date in the past.  We could also expand the usage of the tool by providing a web based user interface, to allow a larger audience to use the tool, including those users that are not comfortable using the command line.  By providing a user friendly web UI, we would also be helping users that ARE comfortable with the command line, by streamlining that usage.  Right now, the tool requires the user to have Goland installed, Go variables set, the project cloned or copied, and dependencies downloaded before they can even install it.  These steps may seem trivial if someone is already using Golang and has some of those steps done, but would mean there is a barrier to entry for using our tool, and the involved process might detour people from even using TheCityFoodTrucks themselves.  Building that web UI would require time and patience, as we would have to design, implement, and test an entire new portion of the program.  It would also involve business decisions about best color scheme, font, logo, and even how buttons are displayed.  Additionally, we'd have to make sure that the tool is accessible through a wide range of web browsers, and if it uses web browser plugins, that we tell the user so that they may have the best experience possible.

As always, before we begin making new features and enhancements, it would be beneficial to have full code coverage through testing.  That would ensure that we do not break current functionality when implementing anything else. 
