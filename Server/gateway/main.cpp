/** File generate by <hotlala8088@gmail.com> 2015/01/14  
 */

#include "config.h"
#include "gwserv.h"
#include "coredump.h"
void assertPrepare(){}
class SignalHandler : public ACE_Event_Handler
{
public:
	virtual int handle_signal (int sig, siginfo_t *, ucontext_t *)
	{
		// �����¼�ѭ��.
		ACE_Reactor::end_event_loop();
		return 0;
	}
};

int ACE_TMAIN (int argc, ACE_TCHAR *argv[])
{
#ifndef ACE_WIN32
	if (argc>1&&strcmp(argv[1],"-d")==0)
	{
		daemonize();
	}
	signal(SIGPIPE,SIG_IGN);
#endif
	Logger::instance()->init();
	Logger::instance()->enableFileOut(".gateway.log","log");
#ifdef ACE_WIN32
	ACE_REACTOR_MAKE;
#elif defined (ACE_HAS_EVENT_POLL) || defined (ACE_HAS_DEV_POLL)
	//ACE_Reactor::instance();
	ACE_Reactor::instance(new ACE_Reactor(new ACE_Dev_Poll_Reactor(32000), 1), 1);
#else
	ACE_Reactor::instance();
#endif
	SignalHandler sh;
	ACE_Reactor::instance()->register_handler(SIGINT, &sh);
	GatewayServ::instance()->reactor(ACE_Reactor::instance());
	GatewayServ::instance()->init(argc,argv);
	ACE_Reactor::run_event_loop();
	return 0;
}