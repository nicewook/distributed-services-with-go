package log

import (
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

// 인덱스 항목 내의 바이트 개수
var (
	offWidth uint64 = 4 // 레코드의 오프셋 정보 uint32 4바이트 - 즉 몇 번째인지
	posWidth uint64 = 8 // 위치(position) 정보 uint64 8바이트 - 즉 정확한 위치
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size()) // 현재 사이즈 저장
	// 일단 최대 사이즈로 Truncate() 해줘서 mmap 대비
	if err = os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}
	if idx.mmap, err = gommap.Map(idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE, gommap.MAP_SHARED); err != nil {
		return nil, err
	}
	return idx, nil
}

func (i *index) Close() error {
	// 메모리 맵 파일부터 싱크
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	// 그 다음 파일 싱크
	if err := i.file.Sync(); err != nil {
		return err
	}
	// 이제 실제 크기만큼 다시 자르기
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

// in 번째 인덱스를 읽어서
// 앞에 4바이트는 out, 그 다음 8바이트는 pos 정보로 파싱해서 리턴하가.
func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}
	if in == -1 {
		out = uint32((i.size / entWidth) - 1) // 가장 마지막 인덱스 계산
	} else {
		out = uint32(in)
	}
	pos = uint64(out) * entWidth // 몇 번째 바이트를 읽을 지 계산
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}
	// out과 pos를 여러번 재활용해서 좀 헷갈린다.
	out = enc.Uint32(i.mmap[pos : pos+offWidth])          // 4바이트 읽기
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth]) // 8바이트 읽기
	return out, pos, nil
}

func (i *index) Write(off uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+entWidth { // 인덱스 하나 추가해도 크기 괜찮은가?
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += uint64(entWidth)
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
