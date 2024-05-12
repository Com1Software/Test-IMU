#include "hbcom.ch"
#include "inkey.ch"

FUNCTION Main( )

   LOCAL cString := ""
   LOCAL nTimeOut := 500
   LOCAL nResult
   
   local line:=""
   LOCAL nPort := 0
   local pa:=DetectSerialPorts()
   local ik:=0

   ?os()
   if len(pa)>0 
      nport=pa[1]
      IF ! hb_comOpen( nPort )
         ? "Cannot open port:", nPort, hb_comGetDevice( nPort ), ;
            "error: " + hb_ntos( hb_comGetError( nPort ) )
      ELSE
         ? "port:", hb_comGetDevice( nPort ), "opened"
         IF ! hb_comInit( nPort, 115200, "N", 8, 1 )
            ? "Cannot initialize port to: 9600:N:8:1", ;
               "error: " + hb_ntos( hb_comGetError( nPort ) )
         ELSE
         
            do while .t.
               ik=inkey()
               if ik = 27 
                  exit
               endif       
               cString := Space( 1 )
               nTimeOut := 500 // 500 milliseconds = 0.5 sec.
               nResult := hb_comRecv( nPort, @cString, hb_BLen( cString ),nTimeOut )
                  IF nResult == 1
               
                  if asc(cstring)=85
                      ?line
                      line=""
                  else
                     line=line+str(asc(cstring))+" " 
                  endif

               ENDIF
            enddo
         ENDIF
         ? "CLOSE:", hb_comClose( nPort )
      ENDIF
   else
      ?"No Serial Ports Found"
   endif
RETURN

FUNCTION DetectSerialPorts()
   local pa:=array(0)
   local x:=25
   local nPort
   local cPortName:="/dev/ttyACM"
   local cPort:=""
   DO WHILE x > 0
      nPort=x
      IF hb_comOpen( nPort )
         aadd(pa,x)
         hb_comClose( nPort )
      ELSE
         cPort=cPortName+alltrim(str(x-1)) 
         hb_comSetDevice( nPort, cPort)
         IF hb_comOpen( nPort )
            aadd(pa,x)
            hb_comClose( nPort )
         ENDIF

      ENDIF
      x--
   enddo
return (pa)
