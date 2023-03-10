# School Manager

This CRUD API is written in Go and manages data to use in a school (students, professor, courses, users).

## Models

- User
  - Schema

```
{
  ID (uint32): unique identifier,
  Name: {
    First (string): first name,
    Middle (string): middle name,
    Last (string): last name
  },
  Email: (string),
  Password: (string),
  UpdatedAt (DateTime),
  CreatedAt (DateTime)
}
```

- Student
  - Schema

```
{
  ID (uint32): unique identifier,
  Name: {
    First (string): first name,
    Middle (string): middle name,
    Last (string): last name
  },
  Address: {
    Raw (string): full address,
    StreetNumber (string): number of the building on street,
    StreetName (string): name of the street,
    Unit (string): unit number of apartment if applicable,
    City (string),
    State (string),
    Country (string),
    CountryCode (string),
    PostalCode (string)
  },
  Courses (Course[]): List of courses currently enrolled,
  Email (string),
  Phone (string),
  GPA (float),
  Credits (float): total number of credits passed,
  AttemptedCredits (float): total number of credits attempted,
  DegreeLevel (string),
  FieldOfStudy (string),
  UpdatedAt (DateTime),
  CreatedAt (DateTime)
}
```

- Professor
  - Schema

```
{
  ID (uint32): unique identifier,
  Name: {
    First (string): first name,
    Middle (string): middle name,
    Last (string): last name
  },
  Address: {
    Raw (string): full address,
    StreetNumber (string): number of the building on street,
    StreetName (string): name of the street,
    Unit (string): unit number of apartment if applicable,
    City (string),
    State (string),
    Country (string),
    CountryCode (string),
    PostalCode (string)
  },
  Email (string),
  Phone (string),
  Courses (Course[]): Courses teaching,
  Position (string)
  UpdatedAt (DateTime),
  CreatedAt (DateTime)
}
```

- Course
  - Schema

```
{
  ID (uint32): unique identifier,
  Department (string),
  CourseNumber (string),
  MaxCapacity (number): number of students allowed in class,
  Professors (uint32[]): Ids of professors teaching this course,
  Students (uint32[]): Ids of students enrolled in course,
  UpdatedAt (DateTime),
  CreatedAt (DateTime)
}
```

## Set Up

### Postgres setup

- Create databases
  - course_manager
  - course_manger_test
- Upon startup will create schema and tables necessary for that school.

### Env File

Create `.env` file such as the following

```
# Postgres Live
API_SECRET=
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgrespw
DB_NAME=course_manager
DB_PORT=32768

# Postgres Test
TestApiSecret=
TestDbHost=localhost
TestDbUser=postgres
TestDbPassword=postgrespw
TestDbName=course_manager_test
TestDbPort=32768
```

API secrets can be anything. DB user and password is what you made your user for the database.

### Run the program

`go run main.go`

### Authentication

Hit endpoint `/login` (POST) with email and password

```json
{
  "email": "jamestest@example.com",
  "password": "password"
}
```

to retrieve Auth Token and set header on further requests with key "Authorization" and value "Bearer <AUTH TOKEN>"

## Entity Endpoints

### Users

#### Get All Users (GET)

Gets all users
`/users`
Returns all users from query.

#### Get User (GET)

Returns a specified user given the id
`/users/:user_id`
Returns that specific user

#### Update User (POST)

Updates only the fields given for a specified user given the id. Updates given fields
`/users/:user_id`
Example Body

```
{
  "name": {
    "first": "Mickey",
    "middle": "disney",
    "last": "Mouse"
  },
  "email": "mickey@example.com",
  "password": "password"
}
```

returns user updated

#### Create User (POST)

Creates new User
`/users`
Example Body

```
{
  name: {
    first (string): "Mickey",
    middle (string): "Disney",
    last (string): "Mouse"
  },
  email: "mickeytest@example.com",
  password: "password"
}
```

returns user created

#### Delete User (DELETE)

Deletes User
`/users/:user_id`
returns ID of user deleted

### Student

#### Get All Student (GET)

Gets all students
`/students`
Returns all students.

#### Get Student (GET)

`/students/:student_id`
Returns all data of student

#### Update Student (POST)

Updates only the fields given for a specified user given the id
`/students/:student_id`
Example Body

```
{
  "name": {
    "first": "Mickey",
    "middle": "Disney",
    "last": "Mouse"
  },
  Address: {
    raw: "136 Highland Dr Burkburnett, TX, 76354",
    street_number: "136",
    street_name: "Highland Dr",
    city: "Burkburnett",
    state: "Texas",
    country: "United States of America",
    country_code: "US",
    postal_code: "76354"
  },
  email: "mickeytest@example.com",
  phone: "(940) 569-3810",
  gpa: 3.15,
  credits: 20,
  attempted_credits: 20,
  degree_level: "bachelors",
  field_of_study: "Mathematics"
}
```

returns updated students

#### Create Student (POST)

Create new Student
`/students`
Example Body

```
{
  "name": {
    "first": "Mickey",
    "middle": "Disney",
    "last": "Mouse"
  },
  Address: {
    raw: "136 Highland Dr Burkburnett, TX, 76354",
    street_number: "136",
    street_name: "Highland Dr",
    city: "Burkburnett",
    state: "Texas",
    country: "United States of America",
    country_code: "US",
    postal_code: "76354"
  },
  email: "mickeytest@example.com",
  phone: "(940) 569-3810",
  gpa: 3.15,
  credits: 20,
  attempted_credits: 20,
  degree_level: "bachelors",
  field_of_study: "Mathematics"
}
```

Returns newly created student

#### Delete Student (DELETE)

Deletes Student
`/students/:student_id`
returns ID of student deleted

### Professor

#### Get All Professors

Gets all Professors
`/professor`
Returns all professors in the query.

#### Get Professor (GET)

Returns a specified user given the id
`/professor/:professor_id`
Returns all data for given professor

#### Update Professor (POST)

Updates only the fields given for a specified faculty given the id
`/professor/:professor_id`
Body

- Allowed fields

Example Body

```json
{
  "name": {
    "first": "Donald",
    "middle": "Disney",
    "last": "Duck"
  },
  "address": {
    raw: "136 Highland Dr Burkburnett, TX, 76354",
    street_number: "136",
    street_name: "Highland Dr",
    city: "Burkburnett",
    state: "Texas",
    country: "United States of America",
    country_code: "US",
    postal_code: "76354"
  },
  "email": "donaldtest@example.com",
  "phone": "(940) 569-3810",
  "education" [
    {
      "school": "University of Illinois",
      "graduation_date": "2019-12",
      "gpa": 3.6,
      "degree_level": "Doctorates",
      "field_of_study: "Computer Science"
    }
  ],
  "office_hour": [
    {
      "start": ,
      "end": ,
      "room":
    }
  ],
  "courses" (uint32[]): [123, 456],
  "position": "Associate Professor"
}
```

Returns professor updated

#### Create Professor (POST)

Creates a new Professor
`/professor/:professor_id`
Example Body

```json
{
  "name": {
    "first": "Donald",
    "middle": "Disney",
    "last": "Duck"
  },
  "address": {
    "raw": "136 Highland Dr Burkburnett, TX, 76354",
    "street_number": "136",
    "street_name": "Highland Dr",
    "city": "Burkburnett",
    "state": "Texas",
    "country": "United States of America",
    "country_code": "US",
    "postal_code": "76354"
  },
  "email": "donaldtest@example.com",
  "phone": "(940) 569-3810",
  "education" [
    {
      "school": "University of Illinois",
      "graduation_date": "2019-12",
      "gpa": 3.6,
      "degree_level": "Doctorates",
      "field_of_study: "Computer Science"
    }
  ],
  "courses": [],
  "position": "Associate Professor"
}
```

Returns newly created professor

#### Delete Professor (DELETE)

Deletes Professor
`/professor/:professor_id`
returns ID of faculty deleted

### Course

#### Get all courses (GET)

Gets all courses
`/course`
Parameters
Returns all courses from query

#### Get Course (GET)

Gets all data for a course
`/course/:course_id`
Parameters

- select (string): SQL style select statement for fields to return '\*' not allowedh
  Returns all data for course

#### Update Course (POST)

Updates only the fields given for a specified course given the id
`/course/:course_id`

Example Body

```json
{
  college: ,
  department: ,
  course_number: ,
  max_capacity: ,
}
```

Returns course updated

#### Create Course (POST)

Creates new Course
`/course`
Example Body

```json
{
  department: ,
  course_number: ,
  max_capacity: ,
  session_times [
    {
      start: ,
      end: ,
      room:
    }
  ],
  professors: [],
  students: []
}
```

returns course created

#### Enroll Student (PUT)

Enrolls a student to a course
`/course/:course_id/student/:student_id/enroll`
Returns field `students` on Course

#### Remove Student (PUT)

Enrolls a student to a course
`/course/:course_id/student/:student_id/remove`
Returns field `students` on Course

#### Assign Professor (PUT)

Assigns a professor to teach course
`/course/:course_id/professor/:professor_id/assign`
Returns field `professors` on Course

#### Remove Professor (PUT)

Removes professor from course
`/course/:course_id/professor/:professor_id/remove`
Returns field `professors` on Course

#### Delete Course (DELETE)

Deletes Course
`/course/:course_id`
returns ID of course deleted
