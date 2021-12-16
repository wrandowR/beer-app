package main


func main(){

	
func init() {
//
}

func main() {

	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	/*

	appController := controller.AppController{
		Organization:    controller.OrgController,
		UserPermissions: controller.UserPermController,
		User:            controller.UserController,
		Project:         controller.ProController,
		Did:             controller.DiController,
		SipTrunk:        controller.SipTController,
		Email:           controller.Email,
		Opa:             controller.OPAController,
		BotWhatsappLine: controller.BotWhatsappLine,
		Bot:             controller.Bot,
		Billing:         controller.BillingCtrl,
		Sms:             controller.SMSController,
		Whatsapp:        controller.Whatsapp,
		Dashboard:       controller.Dashboard,
	}
	*/

	e := echo.New()
	

	e.HTTPErrorHandler = clog.ErrHandler(logrus.StandardLogger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())	
	e.Use(middleware.CORS())
	

	e.GET("/healthz", healthHandler)

	e = router.NewRouter(e, appController)

	go func() {
		if err := e.Start(config.HTTPListener()); err != nil {
			logrus.WithError(err).Error("shutting down server")
		}
	}()
	

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		clog.Panic(log, merry.Wrap(err))
	}
}

func healthHandler(c echo.Context) error {
	row := datastore.SQLDB.QueryRow("SELECT 1")
	var val int
	if err := row.Scan(&val); err != nil {
		return merry.Wrap(err)
	}
	if val != 1 {
		return merry.New("error query pg")
	}
	return c.NoContent(http.StatusOK)
}