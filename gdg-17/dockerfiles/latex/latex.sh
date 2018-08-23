#!/bin/bash
docker container run -it --rm \
	-v $PWD:/data \ 	# Mount Host $PWD to Container /data // HL
	-h latex \		# hostname 
	-u  1000 \		# uid      
	--name latex \
	mfrw/latex
