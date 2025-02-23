package gameplay

// commit, revert, broadcast, publicify actions for game

func CommitGameAction(args *HandlersArgs) error {
	return nil
}

func AbortGameAction(args *HandlersArgs) error {
	defer args.ctx.Done()

	return nil
}

func BroadcastGameAction(args *HandlersArgs) error {
	return nil
}

func PublicDataGameAction(args *HandlersArgs) error {
	return nil
}

func GetGameAction(args *HandlersArgs) error {
	return nil
}
