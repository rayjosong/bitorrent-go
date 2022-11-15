package main

import (
	"io"
	"net/url"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte   // split from pieces
	PieceHashes [][20]byte // split from pieces
	PieceLength int
	Length      int
	Name        string
}

func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error) {
	// TODO: logic to convert toTorrentFile struct
}

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, err) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}

	params := url.Values{}
}
