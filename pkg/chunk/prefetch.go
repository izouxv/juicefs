/*
 * JuiceFS, Copyright 2020 Juicedata, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chunk

import (
	"sync"

	"github.com/juicedata/juicefs/pkg/meta"
)

type keyIno struct {
	Key string
	ino meta.Ino
}
type prefetcher struct {
	sync.Mutex
	pending chan keyIno
	busy    map[string]bool
	op      func(key string, ino meta.Ino)
}

func newPrefetcher(parallel int, fetch func(string, meta.Ino)) *prefetcher {
	p := &prefetcher{
		pending: make(chan keyIno, 10),
		busy:    make(map[string]bool),
		op:      fetch,
	}
	for i := 0; i < parallel; i++ {
		go p.do()
	}
	return p
}

func (p *prefetcher) do() {
	for key := range p.pending {
		p.Lock()
		if _, ok := p.busy[key.Key]; !ok {
			p.busy[key.Key] = true
			p.Unlock()

			p.op(key.Key, key.ino)

			p.Lock()
			delete(p.busy, key.Key)
		}
		p.Unlock()
	}
}

func (p *prefetcher) fetch(key string, ino meta.Ino) {
	select {
	case p.pending <- keyIno{key, ino}:
	default:
	}
}
