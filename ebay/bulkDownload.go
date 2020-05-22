package ebay

func (ec *EbayClient) DownloadAllResponses(itemIDs []string, targetDirectory string) <- chan *error {

	return ec.downloadResp(ec.getItemRawResponse(itemIDs...), targetDirectory)
}
