# Assignment 4 - CSE6690 - Software Development & Operations
tl;dr: Set up Xenial LTS on Vagrant+Virtualbox. Write a tiny web app that interacts with Redis.

See the wiki for the [full requirements](https://github.com/eenblam/devops-proj-4/wiki).

Screencast of demo [here](https://asciinema.org/a/psBuGLWEJVApniFm2aTige3G0).

## Conclusion
I haven't configured NGINX in a while, so I went with a really barebones configuration. I had a lot of fun figuring stuff out with Go. The dependency injection pattern for managing the Redis client was nice, as it's the sort of thing I would attempt in a very basic Python application. I didn't bother with any explicit concurrency given the scope of the application. Provisioning everything was pretty straightfoward, but I didn't have time to brush back up on Ansible. Instead, I just crammed most of that into `provision.sh`, which isn't at all idempotent.

It should be straightforward to clone this repo and replicate the linked demo. I messed something up with my system.d unit, though, so you might have to start Redis manually.

*Note: `provision.sh` includes comments for guides followed in writing it.*
