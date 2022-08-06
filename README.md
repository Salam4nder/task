# Todo CLI

>Simple Todo app that lives in your terminal.

https://user-images.githubusercontent.com/89915813/183265571-657e3b76-a7e1-471c-a490-a360d1906489.mov


# Install

```zsh 
go install github.com/Salam4nder/TodoCli
```

# Usage

## Commands

<img width="1000" alt="Screen Shot 2022-08-06 at 22 52 43 PM" src="https://user-images.githubusercontent.com/89915813/183265683-749a2082-4cf6-4415-84a1-e621e20fadfb.png">

```
todo add <task> // Adds a task to the list
todo ls // List all tasks in the list
todo done <number> // Marks a task with the given index as done
todo ref // Refreshes the list by removing tasks that are marked done
todo rm <number> // Removes task with given index from the list
todo clear // Removes all tasks from the list

todo add --help // if you need additional info
```

# Todo (haha):

* Figure out a way to be able to add todos with special characters. For example " ' "
* Make the UI fancier?
