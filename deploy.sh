# ARC (Azure Container Registry)
# uname: user name generated for the ACR
# upassword: user password generated for the ACR
uname=""
upassword=""

# Regex
regex="^[A-Za-z?_-]+$"
arcregex="^[A-Za-z0-9?.]+$"

out () {
    echo >&2 "$@"
    exit 1
}

check(){
  [[ "$1" =~ $2 ]] || out "Invalid value, $1 provided"
  echo $1 | grep -Eq "$2" || out "Alphanumeric argument required, $1 provided"
}

[ "$#" -eq 2 ] || out "1 argument required, $# provided"
check $1 $regex
check $2 $arcregex

echo "App name: $1"
echo "ACR: $2"
image="$2.azurecr.io/$1:latest"
build_step="docker build --rm -f "Dockerfile" -t "$image" ."

if $build_step; then
  echo "Image Builded Successfully!: $build_step"
  echo "\nImage deploy: docker push $image\n\n"
else
  echo "Error building image $build_step"
fi

read -p "Are you sure to deploy the image? (Y/y) - Press any key to cancel:" -n 1 -r
echo    # (optional) move to a new line
if [[ ! $REPLY =~ ^[Yy]$ ]]
then
    echo "Deployment cancelled"
    exit 0
fi

if az acr login --name $2; then
  echo "Logged!"
else
  echo "Error logging on arc!"
fi

docker login "$2.azurecr.io" -u "$uname" --password "$upassword"

if docker push $image; then
  echo "\nDeployment Finished!"
else
  echo "\nError deploying image: $image"
fi
