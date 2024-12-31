package nxml

type Page struct {
	Item                string   `xml:"item,attr"`
	ScreenID            string   `xml:"ScreenID"`
	TabText             string   `xml:"TabText"`
	TabTextColorR       string   `xml:"TabTextColor>R"`
	TabTextColorG       string   `xml:"TabTextColor>G"`
	TabTextColorB       string   `xml:"TabTextColor>B"`
	TabTextActiveColorR string   `xml:"TabTextActiveColor>R"`
	TabTextActiveColorG string   `xml:"TabTextActiveColor>G"`
	TabTextActiveColorB string   `xml:"TabTextActiveColor>B"`
	DrawTemplate        string   `xml:"DrawTemplate"`
	Pieces              []string `xml:"Pieces"`
	LocationX           string   `xml:"Location>X"`
	LocationY           string   `xml:"Location>Y"`
	SizeX               string   `xml:"Size>CX"`
	SizeY               string   `xml:"Size>CY"`
}

/*
 <Page item="IDW_ItemDescriptionTab">
    <ScreenID>ItemDescriptionTab</ScreenID>
    <TabText>Description</TabText>
    <TabTextColor>
      <R>255</R>
      <G>255</G>
      <B>255</B>
    </TabTextColor>
    <TabTextActiveColor>
      <R>255</R>
      <G>255</G>
      <B>0</B>
    </TabTextActiveColor>
    <DrawTemplate>WDT_Def</DrawTemplate>
    <Pieces>IDW_MadeBy</Pieces>
    <Pieces>IDW_ItemDescription</Pieces>
    <Pieces>IDW_IconButton</Pieces>
    <Pieces>Screen:IDW_Appearance_Socket_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_1_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_2_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_3_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_4_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_5_Screen</Pieces>
    <Pieces>Screen:IDW_Socket_Slot_6_Screen</Pieces>
    <Pieces>IDW_ModButton</Pieces>
	<Pieces>IDW_ItemInfo1</Pieces>
	<Pieces>IDW_ItemInfo2</Pieces>
	<Pieces>IDW_ItemInfo3</Pieces>
	<Pieces>IDW_ItemInfo4</Pieces>
	<Pieces>IDW_ItemInfo5</Pieces>
	<Pieces>IDW_ItemInfo6</Pieces>
	<Pieces>IDW_ItemInfo7</Pieces>
	<Pieces>IDW_ItemInfo8</Pieces>
	<Pieces>IDW_ItemInfo9</Pieces>
	<Pieces>IDW_ItemInfo10</Pieces>
    <Pieces>IDW_ModButtonLabel</Pieces>
    <Pieces>IDW_Row1Col1Stat</Pieces>
    <Pieces>IDW_Row2Col1Stat</Pieces>
    <Pieces>IDW_Row3Col1Stat</Pieces>
    <Pieces>IDW_Row4Col1Stat</Pieces>
    <Pieces>IDW_Row5Col1Stat</Pieces>
    <Pieces>IDW_Row6Col1Stat</Pieces>
    <Pieces>IDW_Row7Col1Stat</Pieces>
    <Pieces>IDW_Row8Col1Stat</Pieces>
    <Pieces>IDW_Row9Col1Stat</Pieces>
    <Pieces>IDW_Row10Col1Stat</Pieces>
    <Pieces>IDW_Row11Col1Stat</Pieces>
    <Pieces>IDW_Row12Col1Stat</Pieces>
    <Pieces>IDW_Row13Col1Stat</Pieces>
    <Pieces>IDW_Row14Col1Stat</Pieces>
    <Pieces>IDW_Row15Col1Stat</Pieces>
    <Pieces>IDW_Row16Col1Stat</Pieces>
    <Pieces>IDW_Row17Col1Stat</Pieces>
    <Pieces>IDW_Row18Col1Stat</Pieces>
    <Pieces>IDW_Row19Col1Stat</Pieces>
    <Pieces>IDW_Row20Col1Stat</Pieces>
    <Pieces>IDW_Row21Col1Stat</Pieces>
    <Pieces>IDW_Row22Col1Stat</Pieces>
    <Pieces>IDW_Row23Col1Stat</Pieces>
    <Pieces>IDW_Row24Col1Stat</Pieces>
    <Pieces>IDW_Row25Col1Stat</Pieces>
    <Pieces>IDW_Row26Col1Stat</Pieces>


    <Pieces>IDW_Row1Col1Value</Pieces>
    <Pieces>IDW_Row2Col1Value</Pieces>
    <Pieces>IDW_Row3Col1Value</Pieces>
    <Pieces>IDW_Row4Col1Value</Pieces>
    <Pieces>IDW_Row5Col1Value</Pieces>
    <Pieces>IDW_Row6Col1Value</Pieces>
    <Pieces>IDW_Row7Col1Value</Pieces>
    <Pieces>IDW_Row8Col1Value</Pieces>
    <Pieces>IDW_Row9Col1Value</Pieces>
    <Pieces>IDW_Row10Col1Value</Pieces>
    <Pieces>IDW_Row11Col1Value</Pieces>
    <Pieces>IDW_Row12Col1Value</Pieces>
    <Pieces>IDW_Row13Col1Value</Pieces>
    <Pieces>IDW_Row14Col1Value</Pieces>
    <Pieces>IDW_Row15Col1Value</Pieces>
    <Pieces>IDW_Row16Col1Value</Pieces>
    <Pieces>IDW_Row17Col1Value</Pieces>
    <Pieces>IDW_Row18Col1Value</Pieces>
    <Pieces>IDW_Row19Col1Value</Pieces>
    <Pieces>IDW_Row20Col1Value</Pieces>
    <Pieces>IDW_Row21Col1Value</Pieces>
    <Pieces>IDW_Row22Col1Value</Pieces>
    <Pieces>IDW_Row23Col1Value</Pieces>
    <Pieces>IDW_Row24Col1Value</Pieces>
    <Pieces>IDW_Row25Col1Value</Pieces>
    <Pieces>IDW_Row26Col1Value</Pieces>

    <Pieces>IDW_Row1Col2Stat</Pieces>
    <Pieces>IDW_Row2Col2Stat</Pieces>
    <Pieces>IDW_Row3Col2Stat</Pieces>
    <Pieces>IDW_Row4Col2Stat</Pieces>
    <Pieces>IDW_Row5Col2Stat</Pieces>
    <Pieces>IDW_Row6Col2Stat</Pieces>
    <Pieces>IDW_Row7Col2Stat</Pieces>
    <Pieces>IDW_Row8Col2Stat</Pieces>
    <Pieces>IDW_Row9Col2Stat</Pieces>
    <Pieces>IDW_Row10Col2Stat</Pieces>
    <Pieces>IDW_Row11Col2Stat</Pieces>
    <Pieces>IDW_Row12Col2Stat</Pieces>
    <Pieces>IDW_Row13Col2Stat</Pieces>
    <Pieces>IDW_Row14Col2Stat</Pieces>
    <Pieces>IDW_Row15Col2Stat</Pieces>
    <Pieces>IDW_Row16Col2Stat</Pieces>
    <Pieces>IDW_Row17Col2Stat</Pieces>
    <Pieces>IDW_Row18Col2Stat</Pieces>
    <Pieces>IDW_Row19Col2Stat</Pieces>
    <Pieces>IDW_Row20Col2Stat</Pieces>
    <Pieces>IDW_Row21Col2Stat</Pieces>
    <Pieces>IDW_Row22Col2Stat</Pieces>
    <Pieces>IDW_Row23Col2Stat</Pieces>
    <Pieces>IDW_Row24Col2Stat</Pieces>
    <Pieces>IDW_Row25Col2Stat</Pieces>
    <Pieces>IDW_Row26Col2Stat</Pieces>


    <Pieces>IDW_Row1Col2Value</Pieces>
    <Pieces>IDW_Row2Col2Value</Pieces>
    <Pieces>IDW_Row3Col2Value</Pieces>
    <Pieces>IDW_Row4Col2Value</Pieces>
    <Pieces>IDW_Row5Col2Value</Pieces>
    <Pieces>IDW_Row6Col2Value</Pieces>
    <Pieces>IDW_Row7Col2Value</Pieces>
    <Pieces>IDW_Row8Col2Value</Pieces>
    <Pieces>IDW_Row9Col2Value</Pieces>
    <Pieces>IDW_Row10Col2Value</Pieces>
    <Pieces>IDW_Row11Col2Value</Pieces>
    <Pieces>IDW_Row12Col2Value</Pieces>
    <Pieces>IDW_Row13Col2Value</Pieces>
    <Pieces>IDW_Row14Col2Value</Pieces>
    <Pieces>IDW_Row15Col2Value</Pieces>
    <Pieces>IDW_Row16Col2Value</Pieces>
    <Pieces>IDW_Row17Col2Value</Pieces>
    <Pieces>IDW_Row18Col2Value</Pieces>
    <Pieces>IDW_Row19Col2Value</Pieces>
    <Pieces>IDW_Row20Col2Value</Pieces>
    <Pieces>IDW_Row21Col2Value</Pieces>
    <Pieces>IDW_Row22Col2Value</Pieces>
    <Pieces>IDW_Row23Col2Value</Pieces>
    <Pieces>IDW_Row24Col2Value</Pieces>
    <Pieces>IDW_Row25Col2Value</Pieces>
    <Pieces>IDW_Row26Col2Value</Pieces>

    <Pieces>IDW_Row1Col3Stat</Pieces>
    <Pieces>IDW_Row2Col3Stat</Pieces>
    <Pieces>IDW_Row3Col3Stat</Pieces>
    <Pieces>IDW_Row4Col3Stat</Pieces>
    <Pieces>IDW_Row5Col3Stat</Pieces>
    <Pieces>IDW_Row6Col3Stat</Pieces>
    <Pieces>IDW_Row7Col3Stat</Pieces>
    <Pieces>IDW_Row8Col3Stat</Pieces>
    <Pieces>IDW_Row9Col3Stat</Pieces>
    <Pieces>IDW_Row10Col3Stat</Pieces>
    <Pieces>IDW_Row11Col3Stat</Pieces>
    <Pieces>IDW_Row12Col3Stat</Pieces>
    <Pieces>IDW_Row13Col3Stat</Pieces>
    <Pieces>IDW_Row14Col3Stat</Pieces>
    <Pieces>IDW_Row15Col3Stat</Pieces>
    <Pieces>IDW_Row16Col3Stat</Pieces>
    <Pieces>IDW_Row17Col3Stat</Pieces>
    <Pieces>IDW_Row18Col3Stat</Pieces>
    <Pieces>IDW_Row19Col3Stat</Pieces>
    <Pieces>IDW_Row20Col3Stat</Pieces>
    <Pieces>IDW_Row21Col3Stat</Pieces>
    <Pieces>IDW_Row22Col3Stat</Pieces>
    <Pieces>IDW_Row23Col3Stat</Pieces>
    <Pieces>IDW_Row24Col3Stat</Pieces>
    <Pieces>IDW_Row25Col3Stat</Pieces>
    <Pieces>IDW_Row26Col3Stat</Pieces>

    <Pieces>IDW_Row1Col3Value</Pieces>
    <Pieces>IDW_Row2Col3Value</Pieces>
    <Pieces>IDW_Row3Col3Value</Pieces>
    <Pieces>IDW_Row4Col3Value</Pieces>
    <Pieces>IDW_Row5Col3Value</Pieces>
    <Pieces>IDW_Row6Col3Value</Pieces>
    <Pieces>IDW_Row7Col3Value</Pieces>
    <Pieces>IDW_Row8Col3Value</Pieces>
    <Pieces>IDW_Row9Col3Value</Pieces>
    <Pieces>IDW_Row10Col3Value</Pieces>
    <Pieces>IDW_Row11Col3Value</Pieces>
    <Pieces>IDW_Row12Col3Value</Pieces>
    <Pieces>IDW_Row13Col3Value</Pieces>
    <Pieces>IDW_Row14Col3Value</Pieces>
    <Pieces>IDW_Row15Col3Value</Pieces>
    <Pieces>IDW_Row16Col3Value</Pieces>
    <Pieces>IDW_Row17Col3Value</Pieces>
    <Pieces>IDW_Row18Col3Value</Pieces>
    <Pieces>IDW_Row19Col3Value</Pieces>
    <Pieces>IDW_Row20Col3Value</Pieces>
    <Pieces>IDW_Row21Col3Value</Pieces>
    <Pieces>IDW_Row22Col3Value</Pieces>
    <Pieces>IDW_Row23Col3Value</Pieces>
    <Pieces>IDW_Row24Col3Value</Pieces>
    <Pieces>IDW_Row25Col3Value</Pieces>
    <Pieces>IDW_Row26Col3Value</Pieces>



    <Pieces>IDW_Heroic1</Pieces>
    <Pieces>IDW_Heroic2</Pieces>
    <Pieces>IDW_Heroic3</Pieces>
    <Pieces>IDW_Heroic4</Pieces>
    <Pieces>IDW_Heroic5</Pieces>
    <Pieces>IDW_Heroic6</Pieces>
    <Pieces>IDW_Heroic7</Pieces>
    <Pieces>IDW_Heroic8</Pieces>
    <Pieces>IDW_Heroic9</Pieces>
    <Pieces>IDW_Heroic10</Pieces>
    <Pieces>IDW_Heroic11</Pieces>
    <Pieces>IDW_Heroic12</Pieces>
    <Pieces>IDW_Heroic13</Pieces>

    <Location>
      <X>0</X>
      <Y>22</Y>
    </Location>
    <Size>
      <CX>361</CX>
      <CY>177</CY>
    </Size>
  </Page>
*/
