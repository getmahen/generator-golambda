# Terraform lambda deployment

	Prerequisites -  Setup credo-auth AWS profile

To initialize your terraform directory:

	source ./setEnv.sh dev 
	terraform plan
	
	
To change environments you will need to call setEnv.sh again with the new env:

	source ./setEnv.sh {{env (dev/qa/prod)}}
