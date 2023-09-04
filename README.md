# How to Run application

## Comments
- The biggest missing part is making use of GoRoutines. 
With more time given I think I could make it in a way to use Go Routines and channels. Probably efficiency/execution time will be less dramatically.
- Code organisation could have been done more focused with OOP in mind.
- I download and extract Recipe JSON file while building the Docker Image. Maybe there can be a different way to handle it. I wanted it to work out of box, so I used it this way.
- Although I also placed option to run it by providing docker volume `-v` flag. So that you can pass your own json data file.

## Build and run application with docker
- Please make sure docker is running on your host.

```
cd /into/project/directory

docker build -t recipe-app .

docker run -v /Users/ismailatkurt/Code/HelloFresh/Project/inputs.json:/app/inputs.json recipe-app
```

You may replace `/Users/ismailatkurt/Code/HelloFresh/Project/inputs.json` with your inputs file.
Sample file is located in the project directory named as `inputs.json`. You can change the values in this file and run again

`docker run -v /Users/ismailatkurt/Code/HelloFresh/Project/inputs.json:/app/inputs.json recipe-app`

### Using your own Recipe JSON File

**Recipe JSON file** is being downloaded while building the image. However, you can specify a file to overwrite it when running `docker run` command by passing `-v local-path:/app/hf_test_calculation_fixtures.json`

Example:
```
docker run -v /Users/ismailatkurt/Code/HelloFresh/Project/inputs.json:/app/inputs.json -v /Users/ismailatkurt/Code/HelloFresh/hf_test_calculation_fixtures.json:/app/hf_test_calculation_fixtures.json recipe-app
```

## Alternative way of running application (Docker Compose)

I have also prepared a docker-compose file. It essentially binds the volume with host to container.
So that you don't need to pass `inputs.json` file path, just change the contents/values of inputs.json and run

```docker-compose up --build```

Please keep in mind this will use the Recipe JSON file that is given by this [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz)

## Run without docker
- Download the recipe file [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz)
- Untar/Unzip the file
- Make sure name of the file is `hf_test_calculation_fixtures.json`
- Then execute go run

```go run .```

# Running Tests

Run tests without displaying details
```
go test ./...
```

Run tests with full output even for passing package tests
```
go test ./... -v
```

Run tests with coverage details. When below command executed it will open a page on browser then you can evaluate coverage details per file.
```
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## GitHub Actions

I have prepared 3 action jobs. Build, Test and Build & Push Docker Image. Since I don't have access to Repo Settings I can't set docker username and token.

However since the repo is private, I hardcoded my credentials (access_token) in GitHub Actions file :) Once you review the challenge I will revoke access token and I don't have any Billing Plan, so it is not useful anyway.

## Pulling already prepared Image and run

```
docker 
```

---
Recipe Stats Calculator
====

In the given assignment we suggest you to process an automatically generated JSON file with recipe data and calculated some stats.

Instructions
-----

1. Clone this repository.
2. Create a new branch called `dev`.
3. Create a pull request from your `dev` branch to the master branch.
4. Reply to the thread you're having with your recruiter telling them we can start reviewing your code

Given
-----

Json fixtures file with recipe data. Download [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz)

_Important notes_

1. Property value `"delivery"` always has the following format: "{weekday} {h}AM - {h}PM", i.e. "Monday 9AM - 5PM"
2. The number of distinct postcodes is lower than `1M`, one postcode is not longer than `10` chars.
3. The number of distinct recipe names is lower than `2K`, one recipe name is not longer than `100` chars.

Functional Requirements
------

1. Count the number of unique recipe names.
2. Count the number of occurences for each unique recipe name (alphabetically ordered by recipe name).
3. Find the postcode with most delivered recipes.
4. Count the number of deliveries to postcode `10120` that lie within the delivery time between `10AM` and `3PM`, examples _(`12AM` denotes midnight)_:
    - `NO` - `9AM - 2PM`
    - `YES` - `10AM - 2PM`
5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
    - Potato
    - Veggie
    - Mushroom

Non-functional Requirements
--------

1. The application is packaged with [Docker](https://www.docker.com/).
2. Setup scripts are provided.
3. The submission is provided as a `CLI` application.
4. The expected output is rendered to `stdout`. Make sure to render only the final `json`. If you need to print additional info or debug, pipe it to `stderr`.
5. It should be possible to (implementation is up to you):  
   a. provide a custom fixtures file as input  
   b. provide custom recipe names to search by (functional reqs. 5)  
   c. provide custom postcode and time window for search (functional reqs. 4)

Expected output
---------------

Generate a JSON file of the following format:

```json5
{
    "unique_recipe_count": 15,
    "count_per_recipe": [
        {
            "recipe": "Mediterranean Baked Veggies",
            "count": 1
        },
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 3
        }
    ],
    "busiest_postcode": {
        "postcode": "10120",
        "delivery_count": 1000
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "11AM",
        "to": "3PM",
        "delivery_count": 500
    },
    "match_by_name": [
        "Mediterranean Baked Veggies", "Speedy Steak Fajitas", "Tex-Mex Tilapia"
    ]
}
```

Review Criteria
---

We expect that the assignment will not take more than 3 - 4 hours of work. In our judgement we rely on common sense
and do not expect production ready code. We are rather instrested in your problem solving skills and command of the programming language that you chose.

It worth mentioning that we will be testing your submission against different input data sets.

__General criteria from most important to less important__:

1. Functional and non-functional requirements are met.
2. Prefer application efficiency over code organisation complexity.
3. Code is readable and comprehensible. Setup instructions and run instructions are provided.
4. Tests are showcased (_no need to cover everything_).
5. Supporting notes on taken decisions and further clarifications are welcome.

