//go:build gomock || generate

package mocks

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination short_header_sealer.go github.com/MahamShakir/quic-go-patch/noninternal/handshake ShortHeaderSealer"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination short_header_opener.go github.com/MahamShakir/quic-go-patch/noninternal/handshake ShortHeaderOpener"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination long_header_opener.go github.com/MahamShakir/quic-go-patch/noninternal/handshake LongHeaderOpener"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination crypto_setup.go github.com/MahamShakir/quic-go-patch/noninternal/handshake CryptoSetup"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination stream_flow_controller.go github.com/MahamShakir/quic-go-patch/noninternal/flowcontrol StreamFlowController"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mocks -destination congestion.go github.com/MahamShakir/quic-go-patch/noninternal/congestion SendAlgorithmWithDebugInfos"
//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package mockackhandler -destination ackhandler/sent_packet_handler.go github.com/MahamShakir/quic-go-patch/noninternal/ackhandler SentPacketHandler"
