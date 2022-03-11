# terraform-provider-pyexec

this project allow to execute python  scirpts using terraform. 

the key to use this module is to put your python script in the same folder , this could be as example:


	resource "python_exec" "main" {
		script = "pythonscript.py"
	    args = "<ARG1> ....."
	}

in case the if  you want to catch the output , you can use this resource(exec_py):

	  output "blabla" {
	    value =  python_exec.main.exec_py

	  }
