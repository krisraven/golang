package main

type nike struct {
}

func (n *nike) makeShoe()) iShoe {
	return &nikeShoe{
		sshoe: shoe{
			logo: "nike",
			size: 14
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14
		},
	}
}