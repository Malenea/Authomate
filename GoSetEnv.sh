if [ -z "$1" ]; then
	echo "Please input your Goodreads dev key as argument"
else
	export GR_DEVKEY="$1"
	echo "$GR_DEVKEY"
fi