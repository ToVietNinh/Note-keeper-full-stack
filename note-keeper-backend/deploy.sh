aws ecr get-login-password --region ap-northeast-1 --profile develop | docker login --username AWS --password-stdin 830357392604.dkr.ecr.ap-northeast-1.amazonaws.com
docker tag dev-adm:latest 830357392604.dkr.ecr.ap-northeast-1.amazonaws.com/yma-development:latest
docker push 830357392604.dkr.ecr.ap-northeast-1.amazonaws.com/yma-development:latest