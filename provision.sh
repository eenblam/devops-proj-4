echo Installing NGINX
sudo echo "deb http://nginx.org/packages/debian xenial nginx" >> /etc/apt/sources.list
sudo echo "deb-src http://nginx.org/packages/debian xenial nginx" >> /etc/apt/sources.list
sudo apt-get update
sudo apt-get install -y nginx
sudo ufw allow 'Nginx Full'
# https://stackoverflow.com/questions/5009324/node-js-nginx-what-now
sudo mv /vagrant/testnginx /etc/nginx/sites-available/testnginx
cd /etc/nginx/sites-enabled/
ln -s /etc/nginx/sites-available/testnginx testnginx
#TODO init.d
sudo systemctl enable nginx
sudo systemctl nginx restart

echo Installing Redis
# See https://redis.io/topics/quickstart
# https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-redis-on-ubuntu-16-04
apt-get install build-essential tcl
cd /tmp
curl -O http://download.redis.io/redis-stable.tar.gz
tar xzvf redis-stable.tar.gz
cd redis-stable
make
make test && sudo make install

echo Configuring Redis
sudo mkdir /etc/redis
sudo cp /vagrant/redis.conf /etc/redis/

echo Configuring systemd for Redis
sudo mv /vagrant/redis.service /etc/systemd/system/redis.service
sudo systemctl enable redis

echo Configuring Redis user
sudo adduser --system --group --no-create-home redis
sudo mkdir /var/lib/redis
sudo chown redis:redis /var/lib/redis
sudo chmod 770 /var/lib/redis

cd

echo Installing golang
curl -O https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

#TODO Export $GOPATH=wherever in the synced folder
export GOPATH=/vagrant/go

echo Installing dep
export PATH=$PATH:$HOME/go/bin
go get -u github.com/golang/dep/cmd/dep

echo Install golang dependencies
cd /vagrant
go get -u github.com/go-redis/redis
