""""""""""""""""""""""""""""""""""""""""""""
" 修改文件字符集
set encoding=utf-8
" 显示行号
set nu
" 自动语法高亮
syntax on
" 输入内容时就显示搜索结果
set incsearch

" 文本改动时自动载入
set autoread
" 代码补全
set completeopt=preview,menu

" 启动游标尺
set ruler
" 突出显示当前行
set cursorline
" 突出显示当前列
"set cursorcolumn
" 语法高亮
set syntax=on
" 语言设置
set langmenu=zh_CN.UTF-8
" 智能补全
set completeopt=longest,menu
" 显示状态栏
set laststatus=2
" 高亮显示搜索
set hlsearch
" 末行模式显示，状态栏显示命令
set wildmenu

" 主题
set background=dark
"colorscheme solarized
"colorscheme gotham256
"colorscheme monokai
"colorscheme neodark
colorscheme space-vim-dark
"hi Comment cterm=italic
hi Normal cterm=NONE guibg=NONE
hi LineNr cterm=NONE guibg=NONE
hi SignColum cterm=NONE guibg=NONE
let g:space_vim_dark_background = 233


" 解决退格
set backspace=indent,eol,start


" vim-compiler 设置
"autocmd FileType go compiler golang

" 错误检查插件
" syntastic 配置
execute pathogen#infect()
"set statusline+=%#warningmsg#
"set statusline+=%{SyntasticStatuslineFlag()}
"set statusline+=%*

let g:syntastic_check_on_open = 0
"let g:syntastic_always_populate_loc_list = 1
"let g:syntastic_auto_loc_list = 1
"let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
execute pathogen#infect()
set nocompatible
" 是否开启文件类型识别
filetype on

filetype plugin indent on
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" VundleVim管理vim插件
Plugin 'VundleVim/Vundle.vim'

" 安装 YouCompleteMe 
"Plugin 'Valloric/YouCompleteMe', { 'do': './install.py --clang-completer' }
Plugin 'Valloric/YouCompleteMe'
" YouCompleteMe 配置
"let g:ycm_error_symbol = '>>'
"let g:ycm_warning_symbol = '>'
"let g:ycm_register_as_syntastic_checker = 0
"let g:ycm_min_num_of_chars_for_completion = 10
"let g:ycm_min_num_identifier_candidate_chars = 10
"let g:ycm_filetype_whitelist = { 'cpp': 1 }
"let g:ycm_filetype_specific_completion_to_disable = { 'cpp': 1 }
"let g:ycm_cache_omnifunc = 0



" 变量、函数、包标签
Bundle 'majutsushi/tagbar'
let g:tagbar_sort = 0
let g:tagbar_width = 50

" 目录树
Plugin 'scrooloose/nerdtree'
map <F8> :NERDTreeToggle<CR>
" 自动开启目录树
"autocmd VimEnter * NERDTree

Plugin 'Yggdroot/LeaderF', { 'do': './install.sh' }

" 安装 vim-gocode
Bundle 'Blackrush/vim-gocode'
Plugin 'nsf/gocode',{'rtp':'vim/'}
"imap <F6> <C-x><C-o>

" 引用查看
Plugin 'dgryski/vim-godef'
let g:go_fmt_autosave = 0
" 如果目标在本文件中则不开启新的窗口
let g:godef_same_file_in_same_window = 1
let g:godef_split = 2

" 安装vim-go
Plugin 'fatih/vim-go'
  
" go bin
let g:go_bin_path = "$GOBIN"
let g:go_fmt_command = "goimports"
let g:go_metalinter_autosave = 1
let g:go_metalinter_autosave_enabled = ['errcheck']
let g:go_metalinter_deadline = "30s"
let g:go_list_height = 20

let g:go_highlight_extra_types = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_fields = 1
let g:go_highlight_interfaces = 0
let g:go_highlight_structs = 0
let g:go_highlight_operators = 0
let g:go_highlight_build_constraints = 1
let g:go_highlight_format_strings = 1
let g:go_auto_type_info = 0

let g:go_guru_scope = ["maid", "Gout"]

let g:tagbar_type_go = { 
	\ 'ctagstype' : 'go', 
	\ 'kinds' : [ 
	\ 'p:package', 
	\ 'i:imports:1', 
	\ 'c:constants', 
	\ 'v:variables', 
	\ 't:types', 
	\ 'n:interfaces', 
	\ 'w:fields', 
	\ 'e:embedded', 
	\ 'm:methods', 
	\ 'r:constructor', 
	\ 'f:functions' 
	\ ], 
	\ 'sro' : '.', 
	\ 'kind2scope' : { 
	\ 't' : 'ctype', 
	\ 'n' : 'ntype' 
	\ }, 
	\ 'scope2kind' : { 
	\ 'ctype' : 't', 
	\ 'ntype' : 'n' 
	\ }, 
	\ 'ctagsbin' : 'gotags', 
	\ 'ctagsargs' : '-sort -silent' 
	\ }

" 错误检查
Plugin 'w0rp/ale'


" 配色插件
Plugin 'altercation/vim-colors-solarized'
Plugin 'KeitaNakamura/neodark.vim'
"Plugin 'liuchengxu/space-vim-dark'

" 命令行 美化
"Plugin 'vim-airline/vim-airline'
"Plugin 'vim-airline/vim-airline-themes'
call vundle#end()


