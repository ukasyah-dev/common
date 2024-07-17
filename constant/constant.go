package constant

const UserID string = "user-id"
const SessionID string = "session-id"
const SuperAdmin string = "super-admin"

type MutationType int64

const MutationCreated MutationType = 0
const MutationUpdated MutationType = 1
const MutationDeleted MutationType = 2
