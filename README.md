欢迎加入 Web3 动物乐园社区 DC      
https://discord.gg/Kc9Qkqdbcd


半自动 grafun脚本
因为普通evm密钥导入hotwallet不会提供evm地址 看不到gra分数   所以必须要用算法来创建
1. 使用命令 go run gra.go 来生成hot wallet 密钥对 并计算mpc地址(hotwallet的evm地址)

2. 运行命令后会输出 /start  1rq8ZO-YgrvF-PmMMBsvyqZUbEpnCBgksp_btRQjLgw=
3. 直接把这个命令发送到 https://t.me/GraFunBot  这个号就完成了

4. 然后go脚本会输出
5.mpc地址|hot密钥|gra分数|计次
6. 0x177d19282e616B19EC5dE297B3Cd5D709399B5Ac|ca97b6658a6fae78915b166b16ff953baa6ab86463c28c37704026cdff50cbf68c567778c59a769ba708cfc334bc33e8eda067ff6f6573847d78f86c62591e3b|900|1

保存好即可（别想投机，必须单地址单tg号 ，不然新地址会顶掉老地址的分数）

如果你有tg协议号  可以做到全自动
hot mpc地址算法 动物乐园内已开源
