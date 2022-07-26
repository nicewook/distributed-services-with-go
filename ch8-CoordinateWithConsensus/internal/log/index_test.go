package log

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	c := Config{}
	c.Segment.MaxIndexBytes = 1024
	idx, err := newIndex(f, c)
	require.NoError(t, err)
	_, _, err = idx.Read(-1) // -1을 읽으라 하면 에러가 나야 한다.
	require.Error(t, err)
	require.Equal(t, f.Name(), idx.Name()) // index 만들때 파일명이 idx.Name()으로 읽혀야 한다.
	entries := []struct {
		Off uint32
		Pos uint64
	}{
		{Off: 0, Pos: 0},  // Offset 0번의 첫 바이트 위치가 0
		{Off: 1, Pos: 10}, // Offset 1번의 첫 바이트 위치가 10
	}

	for _, want := range entries {
		err = idx.Write(want.Off, want.Pos)
		require.NoError(t, err)

		_, pos, err := idx.Read(int64(want.Off))
		require.NoError(t, err)
		require.Equal(t, want.Pos, pos) // 쓴걸 다시 읽어보고 확인
	}

	// 존재 항목 넘어서 읽으려 하면 에러 나야 한다.
	_, _, err = idx.Read(int64(len(entries))) // 최대 인덱스는 len(entires) -1이다.
	require.Equal(t, io.EOF, err)             // 특히나 에러가 EOF 이어야 한다.
	_ = idx.Close()

	// 파일이 이미 있다면 파일의 정보를 바탕으로 index를 만들어야 한다.
	f, _ = os.OpenFile(f.Name(), os.O_RDWR, 0600)
	idx, err = newIndex(f, c)
	require.NoError(t, err)
	off, pos, err := idx.Read(-1) // 가장 마지막 읽기
	require.NoError(t, err)
	require.Equal(t, uint32(1), off)
	require.Equal(t, entries[1].Pos, pos)
}
