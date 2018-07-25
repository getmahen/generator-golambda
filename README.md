#### Go lambda code generator using Yoeman

### STEPS TO CREATE A GO Lambda Scafolding.
- Clone this repo
- Install Yoeman if it is not already installed on your machine. Following command installs Yoeman glabally.
  `npm install -g yo`

- Once Yoeman is installed, install dependencies listed in package.json by running following command
  `cd <into the folder where this source code is cloned>`
  `npm install`

- Install the generator on your machine by running the following command in the working directory
  `npm link`

- After the generator is successfully installed, to generate Go lambda scafolding, create a new folder somewhere on your machine (not in the current working directory)
  `mkdir -p <logforwarder>`  # where logforwarder is the name of your lambda repo
  `cd <into the logforwarder folder>`
  `yo <name of the generator>`
Note: golambda is the name of this generator. You can change it in package.json --> "name" attribute. LEAVE the "generator-" string as is 
in the name. The "generator-" keyword is reserved and that is how Yoeman determines what generators are installed on the machine