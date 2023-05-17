package keeper

import (
	"context"

	"github.com/armon/go-metrics"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorizedAddresses := []string{
		"ununifi1rf5kknlgy9esdcu5qajmef8lg5fph2775vyc0d",
		"ununifi13evn9guc0wn0lkhax4yfd4ef3dcu09yx4l5tru",
		"ununifi1lf6jxtnf4jxsvvny7jyyghmx5m97v0etfgurvg",
		"ununifi1twgmxcrpq8v4uv3y4m8064nrk8q95sfazqxsmm",
		"ununifi1gexk92dukmw37e8e3ud4yx3rd2xc9rc2ptwqw9",
		"ununifi1tnewpa4vppvfyf8k6g2lra94r5stymxpdjykrx",
		"ununifi1dlptmum4n9y9uhzy5yqe8nk6w774pfgz7j0jyl",
		"ununifi1uh6dul863ps5mn5n625uhr6r0e4krv204nk03v",
		"ununifi1e5ff70d8nywu5v4eclh8qjpskpxv02d0fal8qj",
		"ununifi1qwam5r5xhsz7relg2xwuzds6a64g4aaasuwj7q",
		"ununifi1f0s9c0vjxh0zeejylezk0gfhasrfz44fz9wumu",
		"ununifi1frl5rz54ntw6xl9xk97k98gydep84gr7j3z42g",
		"ununifi18k2h8m9qhg4ae565nf64zhe6ycfvqce93ggrlh",
		"ununifi10zuxh2er7snwvuak23y4fv4ztr4g94z095x8n0",
		"ununifi1fgtufgff5jj7lrc3wv2c9cxs5muyksju89p0wr",
		"ununifi1lkxkuvnz4rl78v3uturlwghgrkxjrga26mn9t5",
		"ununifi1u24sylmyp4qsyk3zps9vsfqadana583pkseswk",
		"ununifi1z4px8vum4jp5drzlpkcg8w7u7sdhk4xp8cqyay",
		"ununifi1wkkweckuzgd4tg6l9zyleszksvahvhn6td0jtv",
		"ununifi1e2q3xtq75f3zrw0gy556juwwz2sv5whf8ycuj4",
		"ununifi1khe6yv4zswaergkrzv0dmq3afcda5fx4jjmf07",
		"ununifi1je4h6032v23syjvlgkerrh0g6ynvelu28m84y9",
		"ununifi1ap5yrzzyf6lfhqyuw53rxyh4e068eusazdhkvc",
		"ununifi1qr5nguktymnxlpcu32287w008fhe8527sy0hul",
		"ununifi1p2mnna9xyrqv2ftgq74gazjdfs96a9zmuf0zdv",
		"ununifi148jhz509cywvn4crrgu87xz83hq3k5eyuunke8",
		"ununifi188fuxdc2cfww99xdugw2m55qh69wennzy9zsqn",
		"ununifi1p5kf5ugsm5qgkur5jpf0ehuwuel3u9apgheeav",
		"ununifi17antaqc3sc8j0fkx438k7taua8vc8qdlk72869",
		"ununifi18ew9vtnqr8tq580wmzd0mmufevkr46xszj932e",
		"ununifi14e6gq974lfsnl5fujt4ewxqfwyq8lvj27h2rvt",
		"ununifi1xtrxcpnpz6af44thcpg8d0qshr00yr074jtufc",
		"ununifi1549959r2387ms5ha6ftmm3jadxcq8t8ctll22k",
		"ununifi1hgdu9amkwwl4w4rru96y7vqa5sl8qda8gw343z",
		"ununifi1a3n08zpm0ed7kt5xtp4362e6r26la9x0h9wzux",
		"ununifi1qq9g6x48kvqeefwgy4jm3hjfgpwj229gr383ra",
		"ununifi1sl683ar0yj5g5cvkaal3f0jj5aejdn7lxqwg2x",
		"ununifi14gvs340wy72hj66uc0h7d5vxspy985rjtm4q45",
		"ununifi15krnl7rsu5wjm48yqx7awqjjzqlgnn4smx2z54",
		"ununifi1fgd0etj2tcl43cnmzpem5r0rd9ulxnx2ed3mqc",
		"ununifi190cad4t68d0amv6peqhwaef4rtuwt26vafzz50",
		"ununifi1ytqttunl98sx0cwq3afw2czw3ukamcqpz28t4c",
		"ununifi1la2sdqznz5xgp30c59xdm6law5wn0l3nl2arqp",
		"ununifi1r58l8s5nlwa8hu6eajgnc8gfazmsqzeqtry569",
		"ununifi1y8430kzjeudf8x0zyvcdgqlzcnwt3zqzedayt9",
		"ununifi1hgh096kd8h2w4c9flm56flutytpvywhs58vmpq",
		"ununifi1anc7svftth5zdtagtqt4flrh027r7x936338dq",
		"ununifi12mafz99a0erwyd0gu9jlaanwvlaf9fs53r38nu",
		"ununifi1hlt2km0e6jakjz20pdtru8tt2kqhut9v9twl4z",
		"ununifi126vac8rng0t4zdwesvnznh5w5rqxdrrtrnvskk",
		"ununifi19g2m9p3zd529wyt0ztusut63xkd3358jer2ev0",
		"ununifi1mpv3a9cxfrrkze64454vmra77xrkhk2qql6uwj",
		"ununifi130rpfseu72la5rlnmmejr6w2hc7qhaqsdrh8am",
		"ununifi1lxjfmrd8xrsxxty7kemrz3rpgswrt8vtsq37nc",
		"ununifi15h3drpk7mf0gn38eh3xnwmn4jrnkvlx7kvgj69",
		"ununifi1r500cehqg5u6fhsaysmhu4cnw5pz3lxcqhgaq7",
		"ununifi1pa29ejcfrylh69pvntrx3va9xej69tnx7re567",
		"ununifi15hggf3c67juhfytwcs55pawatl7t3mgmumr2pl",
		"ununifi1ykqpgtt42srf7kxlrktvztslqyqtmqmltw5l55",
		"ununifi1epv847nxzr487gd2uvrrz9fk00wgmxqf74nfqt",
		"ununifi1wynlx3a6x857k4dflxryxtds6g8al6kyzfw0ul",
		"ununifi1tuq8npsuxfxf03m30we425dtlqjhmqg790yl0p",
		"ununifi1kzshf4x3lu0n8hpelgmsame7eddx7stvex5zs9",
		"ununifi18mnrzrky4alnlcsuu4ax5z7sjhdx4vq6cjw4cd",
		"ununifi1hsmytrahnlpggrf27klm5ew3ejx8yxnfst6eaj",
		"ununifi1qm59fuy2sdxwmz0528v8v60xmhvjmpludxzq8j",
		"ununifi1p4gwl8v9geml756n6nqa0evnhuceya250whu0a",
		"ununifi1r8ud9pu2lyc6sq3lky7cjd8elf0fk2d0l7c235",
		"ununifi1epg4f8mwqd9tes2nt9snyas6quhkgsa66ug4l7",
		"ununifi1alm04z20gygaqv4c8f0hcgz44ckgvkckex8v99",
		"ununifi1v7cl3cvvy9l6gctf6e0uuss6xud8a4auk9vtqk",
		"ununifi1hp9cdkkekdspv8z278sw98vdep96cuspadn99f",
		"ununifi1hmdqsnmlflll9qj528kx3g9v8cgfcmkxsg89aj",
		"ununifi18kjnwyxvyfs4yvyqrayjwknslsjt937n293ppf",
		"ununifi1x3w4pmha4pnk5999y5fq549rxsspycjll5r9k5",
		"ununifi106637q507s79vwx80478mykaa2cm25wcxf924v",
		"ununifi1nhvheayxq578j3jce775qgkuzldpa73rp2rlee",
		"ununifi1gall7cq09l8envnhgvsqwdqm28273pyfl24ay9",
		"ununifi172quh59xd0u32jl7f5yvkxlhfadma56g4zp6q3",
		"ununifi1gxahry3sc8q2dcllu6fpgk6uyx8umulkzugy2x",
		"ununifi1t2atpkr7s5l65uywfdqhpxjnfwdzj9zxyxugs8",
		"ununifi1pc3me30qxpv2afn7sar7hulyuqaa75t5yh7tl2",
	}

	if err := k.IsSendEnabledCoins(ctx, msg.Amount...); err != nil {
		for _, addr := range authorizedAddresses {
			if msg.FromAddress == addr {
				goto allowed
			}
		}
		return nil, err
	}
allowed:

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	to, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	if k.BlockedAddr(to) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.ToAddress)
	}

	err = k.SendCoins(ctx, from, to, msg.Amount)
	if err != nil {
		return nil, err
	}

	defer func() {
		for _, a := range msg.Amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", "send"},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()

	return &types.MsgSendResponse{}, nil
}

func (k msgServer) MultiSend(goCtx context.Context, msg *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// NOTE: totalIn == totalOut should already have been checked
	for _, in := range msg.Inputs {
		if err := k.IsSendEnabledCoins(ctx, in.Coins...); err != nil {
			return nil, err
		}
	}

	for _, out := range msg.Outputs {
		accAddr := sdk.MustAccAddressFromBech32(out.Address)

		if k.BlockedAddr(accAddr) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", out.Address)
		}
	}

	err := k.InputOutputCoins(ctx, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}

	return &types.MsgMultiSendResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (k msgServer) SetSendEnabled(goCtx context.Context, msg *types.MsgSetSendEnabled) (*types.MsgSetSendEnabledResponse, error) {
	if k.GetAuthority() != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if len(msg.SendEnabled) > 0 {
		k.SetAllSendEnabled(ctx, msg.SendEnabled)
	}
	if len(msg.UseDefaultFor) > 0 {
		k.DeleteSendEnabled(ctx, msg.UseDefaultFor...)
	}

	return &types.MsgSetSendEnabledResponse{}, nil
}
