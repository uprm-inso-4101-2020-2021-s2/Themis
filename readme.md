# Themis Decentralized Voting Platform

**Themis** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Types
There are four types used around Themis

|  Type   | Description |
| ------- | ----------- |
| Group   | Encapsulates actions inside, used to separate governance from the rest of the chain. Contains Polls and Accounts. |
| Account | Gives a user's wallet permission to have vouchers and vote on polls. |
| Poll    | Created by group admin, represents a poll with a description of 140 characters. |
| Vote    | Vouchers can be converted to votes in polls. |

## CLI Commands

The fastest way to get into interacting with the chain is by using its included CLI commands.

### Setting information

To get a list of available commands:
```bash
Themisd tx Themis -h
```

Create a new group named ExampleGroup under our wallet cosmos01:
```bash
# [group name]
Themisd tx Themis group create ExampleGroup --from cosmo01
```

Change ExampleGroup's name to GroupExample assuming the group's ID is 1234:
```bash
# [group ID] [new group name]
Themisd tx Themis group set-name 1234 GroupExample --from cosmos01
```

Now you're going to want to make a poll so that voting can take place. Let's say you want to ask what type of cookie is
your group member's favorite, you'd need to come up with a name for the poll, have a description, contain the items to
be voted on and finally a deadline, which would be Dec 27 2021 at 3pm.

To create a new poll:
```bash
# [group ID] [title] [description] [deadline] [options..]
Themisd tx Themis poll create 1234 TopCookies "Which of these cookies do you think is the best?" 1640617200 "Chocolate Chip" "Macadamia Nut" "Plain Dough" --from cosmos01
```

To change that poll's description, assuming the poll's ID is 1234-0 :
```bash
# [poll ID] [new description]
Themisd tx Themis poll set-description 1234-0 "Pick your favorite cookie!" --from cosmos01
```

To extend the deadline to 2022:
```bash
# [poll ID] [new deadline]
Themisd tx Themis poll extend-deadline 1234-0 1672153200 --from cosmos01
```

So far you have a group, and a poll made but no accounts, how will your friends vote on your poll? First you need to
deposit vote vouchers on your friend's wallet. Let's deposit 3 vouchers, don't worry, `give-vote` takes care of account creation if one doesn't
exist already. Let's assume your friend's wallet is cosmos02:
```bash
# [group] [receiver] [vote amount]
Themisd tx Themis account set-vouchers 1234 cosmos02 3 --from cosmos01
```

Now your friend wants to vote, so lets go ahead and vote!
```bash
# [poll] [option]
Themisd tx Themis vote create 1234-0 1 --from cosmos02
```