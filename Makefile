mod:
	go mod download

test:
	go test -v

coverage:
	@percentage=$$(go test -tags excludetest -coverpkg=./... | awk '{ if (index($$2,"%") != 0) { print substr($$2,1,length($$2)-1) } }'); \
  if [ "`echo "$${percentage} < 100.0" | bc`" -eq 0 ]; then \
		echo 'Test coverage is above 100%.'; \
		exit 0; \
	else \
		echo 'Test coverage is below 100%!'; \
		exit 1; \
	fi;
