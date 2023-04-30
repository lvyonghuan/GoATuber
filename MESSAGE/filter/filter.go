package sensitive

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Filter 敏感词过滤器
type Filter struct {
	Trie  *Trie
	noise *regexp.Regexp
}

// New 返回一个敏感词过滤器
func New() *Filter {
	return &Filter{
		Trie:  NewTrie(),
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
}

// UpdateNoisePattern 更新去噪模式
func (filter *Filter) UpdateNoisePattern(pattern string) {
	filter.noise = regexp.MustCompile(pattern)
}

// LoadWordDict 加载敏感词字典
func (filter *Filter) LoadWordDict(path string, method int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return filter.Load(f, method)
}

// LoadNetWordDict 加载网络敏感词字典
func (filter *Filter) LoadNetWordDict(url string) error {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	rsp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return filter.Load(rsp.Body, 0)
}

// Load common method to add words
func (filter *Filter) Load(rd io.Reader, method int) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		if method == 1 {
			words := strings.Split(string(line), ",")
			temp := words[0]
			line = []byte(temp)
		}
		filter.Trie.Add(string(line))
	}

	return nil
}

// Validate 检测字符串是否合法
func (filter *Filter) Validate(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.Trie.Validate(text)
}

// FindAll 找到所有匹配词
func (filter *Filter) FindAll(text string) []string {
	return filter.Trie.FindAll(text)
}

// RemoveNoise 去除空格等噪音
func (filter *Filter) RemoveNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
