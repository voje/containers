IMG="voje/john"

build:
	docker build . -t ${IMG}
	docker push ${IMG}

run:
	docker run \
		--rm \
		-it \
		-v $(shell realpath workdir):/workdir \
		${IMG} \
		/bin/bash

# Put your starting strings in ./workdir/wordlist.txt
generate:
	./wordlist_generator.py ./workdir/wordlist.txt --permutations=swap_two,case > ./workdir/generated_wordlist.txt
