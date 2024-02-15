package domain

// SequenceNumberGenerator Local sequence number generator
type SequenceNumberGenerator interface {
    // Next one-to-one (userId1, userId2); one-to-many (userId, channelId)
    Next(int64, int64) int64
}
