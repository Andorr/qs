# QS CLI
A simple command-line interface written in Go to interact with [QS](https://qs.stud.iie.ntnu.no/) - Dataingeniør's queue-system for delivering assignments.

```diff
- Warning: The CLI might be out of date due to further development and changes to QS
- This project was last updated at 15.02.2020 
```

## Contents
- [Installation](#installation)
- [Usage](#usage)
    - [Authentication](#authentication)
    - [Adding subject](#adding-your-subject)
    - [Add to queue](#add-yourself-to-the-queue)
    - [Add to queue with friends](#add-yourself-to-queue-with-friends)
- [Functionality](#functionality)

## Installation
This CLI requires Go to be installed. Run the command below to install the CLI on your computer.
```
go get github.com/andorr/qs/cmd/qs
```

## Usage
After installing the CLI should be available through the command line as such:
```
qs
```

### Authentication
Right now the CLI does not currently support authentication with username and password. To login 
you have to use the cookie-flag and paste in your cookie. You can get your cookie by logging in on
the main website, opening the developer-tools at chrome and copy-pasting the result of typing __document.cookie__
in the console window.
```
qs login --cookie="<YOUR_COOKIE_HERE>"
```

### Adding your subject
To use the CLI you have to tell the CLI which subjects you have. To do so run the following command:
```
qs config subject add <SUBJECT_NAME> <SUBJECT_ID>

// For example
qs config subject add math 1337

// Listing your subjects
qs config subject list
```
You can retrieve the __subject_id__ from the url on the website.

### Add yourself to the queue
You can add yourself to a queue with the CLI. The CLI will spam the QS-server until the queue opens and add you to the queue.
Run the following command:
```
qs queue add <SUBJECT_NAME> <EXERCISES>

// For example
qs queue add math 3
qs queue add math 3,4,5 // To add multiple exercises 

// To specify room and desk
qs queue add math 3 --desk=9 --room=6
qs queue add math 3 --sleep=1000 // 1 second between every request
qs queue add 1337 3 --id // Using raw subject_id here instead of subject_name
```

To view the queue in real-time run the following command:
```
qs queue list <SUBJECT_NAME>

// For example
qs queue list math
```

### Add yourself to queue with friends
To add yourself with friends you have to register your friends first. You can find the __person_id__
for your friends from the developer-tool when visiting the original website.
```
qs config people add <NAME> <PERSON_ID>

// List all your friends
qs config people list
```

To add you and your friends to a queue, do the following:
```
qs queue add <SUBJECT_NAME> <EXERCISES> --group=<NAME1>,<NAME2>,<NAME3>

// Example
qs queue add math 3,4 --group=sveinung,william
```

If you are already in the queue, you can add your friends by running the command below.
You can find the __queue_element_id__ from the developer-tools when browsing the website.
```
qs queue group <QUEUE_ELEMENT_ID> <NAMES> <EXERCISES>

// Example
qs queue group 101 sveinung,william 3,4
```

## Functionality
Currently there is some important functionality missing from the CLI, for example login with password
and username, viewing available rooms and desks...etc. Unless there are demand for adding such functionality
into the CLI it won't be implemented.

If there is something you think is crucial missing from the CLI do not hesitate to make contact. 