package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TeststoreAppendRead(t *testing.T) {
	f, err := os.CreateTemp("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	s, err := newStore(f)
	require.NoError(t, err)

	// 추가하고 읽기 테스트
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	// 다시 읽어도 같은 파일을 가리키는 store를 다시 가져오는 것 확인
	s, err = newStore(f)
	require.NoError(t, err)
	testRead(t, s)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(write)
		require.NoError(t, err)
		require.Equal(t, pos+n, width*i) // pos에서 n만큼 추가했으니 width*i
	}
}

func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ { // i가 uint64일 필요는 없구나
		read, err := s.Read(pos) // pos 위치를 읽는다.size 읽고 그 사이즈 만큼 읽고
		require.NoError(t, err)
		require.Equal(t, write, read)
		pos += width
	}
}

func testReadAt(t *testing.T, s *store) {
	t.Helper()
	for i, off := uint64(1), int64(0); i < 4; i++ {
		b := make([]byte, lenWidth)
		n, err := s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, lenWidth, n)
		off += int64(n)

		size := enc.Uint64(b)
		b = make([]byte, size)
		n, err = s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, write, b)
		require.Equal(t, int(size), n) // 실제로 사이즈만큼 읽었는가?
		off += int64(n)
	}
}

func TestStoreClose(t *testing.T) {
	f, err := os.CreateTemp("", "store_close_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	require.NoError(t, err)
	_, _, err = s.Append(write)
	require.NoError(t, err)

	_, beforeSize, err := openFile(f.Name())
	require.NoError(t, err)

	err = s.Close() //
	require.NoError(t, err)

	_, afterSize, err := openFile(f.Name())
	require.NoError(t, err)
	require.True(t, afterSize > beforeSize)
	t.Logf("beforeSize %d, afterSzie %d", beforeSize, afterSize)
}

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(
		name,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fi.Size(), nil
}
