CONTRACT := Math.sol
GO_CONTRACT := mathsc.go

compile:
	@solc --bin \
		  --abi \
		  --overwrite \
		  --optimize \
		  -o target $(CONTRACT)
	
	@abigen --sol=$(CONTRACT) \
	        --pkg=main \
			--out=$(GO_CONTRACT)

	@echo Success