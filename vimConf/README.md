# 我的 Vim/NeoVim 配置

首先需要配置 `.vimrc` 或 `init.vim` 写入如下内容

```vimscript
let my_vimrc_path = 'init.vim_path'                                                                                                                                 
let g:vim_home = get(g:,'vim_home',expand(my_vimrc_path))

exec 'source'  g:vim_home.'/init.vim'

set secure

```

然后将 此目录下的 `init.vim` 拷贝至上文的 `ini.vim_path` 目录中
