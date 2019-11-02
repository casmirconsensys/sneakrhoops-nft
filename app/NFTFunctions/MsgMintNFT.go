type MsgMintNFT struct {
    Sender    sdk.AccAddress
    Recipient sdk.AccAddress
    ID        string
    Denom     string
    TokenURI  string
}