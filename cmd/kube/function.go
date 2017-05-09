package main

//type function struct {
//
//	// a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
//	// 4122 for a function
//	uuid uuid.UUID
//
//	// name of the function - used in kubernetes
//	name string
//
//	// tracer will receive trace information of function activity
//	tracer trace.Tracer
//}
//
//// newFunction creates a new function that is ready to go.
//func newFunction() *function {
//	return &function{
//		uuid:   uuid.New(),
//		name:   "lambda-",
//		tracer: trace.Off(),
//	}
//}
//
//func (f function) getUidName() string {
//	uidName := make([]byte, len(f.name)+len(f.uuid.String()))
//	copy(uidName, f.name)
//	copy(uidName[len(f.name):], f.uuid.String())
//	return string(uidName)
//}
//
//// A wrapper around http handler and context
//// - A handler of type function
//// - An embedded field of type *kubeContext
//type functionHandler struct {
//	*kubeContext
//	H func(*kubeContext, http.ResponseWriter, *http.Request) (int, error)
//}
//
//// HTTP error code handler
//func (ch functionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	status, err := ch.H(ch.kubeContext, w, r)
//	if err != nil {
//		log.Printf("HTTP %d: %q", status, err)
//		switch status {
//		case http.StatusNotFound:
//			http.NotFound(w, r)
//		case http.StatusInternalServerError:
//			http.Error(w, http.StatusText(status), status)
//		default:
//			http.Error(w, http.StatusText(status), status)
//		}
//	}
//}
//
//// FunctionIndexHandler provides default action when calling root function path
//func FunctionIndexHandler(k *kubeContext, w http.ResponseWriter, r *http.Request) (int, error) {
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//
//	function := newFunction()
//	log.Printf("Function: %q", function)
//	hostname := r.URL.Query().Get("hostname")
//	if len(hostname) != 0 {
//		log.Printf("Hostname: %s", hostname)
//		io.WriteString(w, hostname) // or
//		w.Write([]byte(hostname))
//	}
//	var op operation
//	op = &deployOperation{
//		image: "p7/microservice",
//		name:  function.getUidName(),
//	}
//	op.Do(k.client)
//
//	// functionHandler has
//	//fmt.Fprintf(w, "FunctionIndexHandler: client is %s", k.client)
//	//json.NewEncoder(w).Encode(p)
//	js, err := json.Marshal(function.getUidName())
//	if err != nil {
//		return http.StatusInternalServerError, err
//	}
//	w.Write(js)
//
//	return 200, nil
//}
