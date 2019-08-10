Vagrant.configure 2 do |config|

    config.vm.box = "ubuntu/bionic64"
    config.vm.network "private_network", ip: "192.168.33.10"
    
    config.vm.provision "shell", inline: <<-SHELL
    wget -O- https://packages.erlang-solutions.com/ubuntu/erlang_solutions.asc | apt-key add -

    echo "deb https://packages.erlang-solutions.com/ubuntu bionic contrib" | sudo tee /etc/apt/sources.list.d/erlang.list


    curl -s https://packagecloud.io/install/repositories/rabbitmq/rabbitmq-server/script.deb.sh | bash

    apt-get install -q -y rabbitmq-server

    service rabbitmq-server stop

    cat >/etc/rabbitmq/rabbitmq.config <<EOF
[
    {rabbit,
    [
        {tcp_listeners,         [5672]},
        {log_levels,            [{connection, info}, {channel, info}]},
        {loopback_users,        []},
        {auth_mechanisms,       ['PLAIN', 'AMQPLAIN']},
        {default_vhost,         <<"/">>},
        {default_user,          <<"guest">>},
        {default_pass,          <<"guest">>},
        {default_permissions,   [<<".*">>, <<".*">>, <<".*">>]},
        {default_user_tags,     [administrator]}
    ]}
].
EOF

    rabbitmq-plugins enable rabbitmq_management

    service rabbitmq-server start
    SHELL
end