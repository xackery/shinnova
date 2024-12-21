package nova

type Composite struct {
	Item     string   `xml:"item,attr"`
	Includes []string `xml:"Include"`
}

/*
<Composite>
        <Include>SIDL.xml</Include>
        <Include>EQLSUI_Animations.xml</Include>
        <Include>EQLSUI_Templates.xml</Include>
		<Include>EQLSUI_EulaWnd.xml</Include>
		<Include>EQLSUI_SOESplashWnd.xml</Include>
		<Include>EQLSUI_ESRBSplashWnd.xml</Include>
		<Include>EQLSUI_MainWnd.xml</Include>
		<Include>EQLSUI_ConnectWnd.xml</Include>
		<Include>EQLSUI_ServerSelectWnd.xml</Include>
		<Include>EQLSUI_ChatWnd.xml</Include>
		<Include>EQLSUI_ChatOptionsDialog.xml</Include>
		<Include>EQLSUI_ColorPickerWnd.xml</Include>
		<Include>EQLSUI_CreditsWnd.xml</Include>
		<Include>EQLSUI_NewsWnd.xml</Include>
		<Include>EQLSUI_HelpWnd.xml</Include>
		<Include>EQLSUI_OKDialog.xml</Include>
		<Include>EQLSUI_YesNoDialog.xml</Include>
		<Include>EQLSUI_PollWnd.xml</Include>
		<Include>EQLSUI_OrderExpansionWnd.xml</Include>
		<Include>EQLSUI_WebOrderWnd.xml</Include>
		<Include>EQLSUI_SplashExWnd.xml</Include>
   </Composite>
*/
