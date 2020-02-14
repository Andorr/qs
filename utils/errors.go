package utils

// General errors
const ConfigError = "Was not able to load config.\nError: %s"

const MissingArguments = "Required arguments are missing!\n\nCommand: %s\n"

const AddToGroupError = "Error occurred when adding as group:\t%s\n"

const Unauthenticated = "You are not authenticated. Please use 'qs auth --cookie <COOKIE>'."

const AlreadyExists = "Key '%s' already exists in %s"

const InvalidArgumentInt = "Invalid argument! %s is not a valid id, it must be an integer."

const UnableToSave = "Was not able to save the changes\nError: %s\n"

const MissingEntity = "Key '%s' does not exist in %s."


// Qs specific
const QueueListError = "Was not able to fetch queue by id %d\nError: %s"

const QueueRemoveError = "Was not able to remove %d from queue\nError: %s"