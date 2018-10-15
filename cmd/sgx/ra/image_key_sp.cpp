// image復号化鍵/暗号化鍵を管理，，，，
// これもenclaveでやらないとダメじゃん まあSPは信頼する前提にしてるしいいか
// やるに越したことはないけど

#include "msgio.h"

int main(int argc, char *argv[])
{
    /* TASK クライアントからのimage暗号化鍵取得要求に答えられるようにする */
    // 別サーバーにしてポートを変えたほうが簡素な気がする
    // clientからのimage暗号化鍵要求だった場合，image暗号化鍵をresponseする
    // process_msg01()でblockingしちゃってるからそこはなんとかしないと
    // if (msgio->read() == request_image_crypto_key) { msgio->send(鍵) break; }

    // 本当は証明書とかでちゃんと署名したほうが良いんだろうなぁ

    // 別にGolangでよくね
    
    MsgIO *msgio;

    // 正直ハードコーディングしたい
    // endpointに ip addr を登録
    // http.POST("endpointURL, ipaddr, port");

    // 色々簡略化
    try {
        msgio = new MsgIO(NULL, port);
    } catch(...) {
        return 1;
    }

    while ( msgio->server_loop() ) {
         
    }
}