#include <iostream>
#include <fstream>
#include <string>
#include "demo.pb.h"

using namespace std;

model::SearchReq *createReq();
int writeReqToFile(string &&fn, model::SearchReq *req);

int main()
{
    auto req = createReq();
    cout << "Req.q: " << req->q() << endl;
    cout << "There are " << req->params_size() << " values in my map" << endl;

    writeReqToFile("fromCxx.bin", req);

    // Optional:  Delete all global objects allocated by libprotobuf.
    google::protobuf::ShutdownProtobufLibrary();
    return 0;
}

int writeReqToFile(string &&fn, model::SearchReq *req)
{
    // Write the req to file.
    fstream output(fn, ios::out | ios::trunc | ios::binary);
    if (!req->SerializeToOstream(&output))
    {
        cerr << "Failed to write req to stream." << endl;
        return -1;
    }
}

model::SearchReq *createReq()
{
    model::SearchReq *req = new model::SearchReq();

    string q = "this is my query from C++";
    req->set_q(q);
    auto params = req->mutable_params();

    (*params)["key0"] = "value0";
    (*params)["key1"] = "value1";
    (*params)["key2"] = "value2";

    return req;
}