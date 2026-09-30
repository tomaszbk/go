class Callback {
public:
	virtual ~Callback() { }
	virtual std::string run() { return "Callback::run"; }
};

class Caller {
private:
	Callback *callback_;
public:
	Caller(): callback_(0) { }
	~Caller() { delCallback(); }
	void delCallback() { delete callback_; callback_ = 0; }
	void setCallback(Callback *cb) { delCallback(); callback_ = cb; }
	std::string call();
};
