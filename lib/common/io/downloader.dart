import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:syncreve/base/base_utils.dart';
import 'package:syncreve/common/grpc/grpc_manager.dart';
import 'package:syncreve/data/app/grpc_file_download_info_data.dart';
import 'package:syncreve/generated/grpc/libsyncreve/protos/file_sync.pbgrpc.dart';

class Downloader {
  static const fileDownloadQueueStatusWaiting = 0;
  static const fileDownloadQueueStatusDownloading = 1;
  static const fileDownloadQueueStatusDone = 2;
  static const fileDownloadQueueStatusError = -1;

  static Future<String> addDownloadTask(
      {required String url,
      required String savePath,
      required String fileName,
      required String cookie,
      required DownloadInfoRequestType type}) async {
    final r = DownloadTaskRequest(
        url: url,
        savePath: savePath,
        fileName: fileName,
        cookie: cookie,
        downLoadType: type);

    if (await File("$savePath/$fileName").exists()) {
      dPrint("[Downloader] addDownloadTask file exists");
      throw "file exists";
    }

    dPrint("[Downloader] addDownloadTask r==$r");

    final id = await AppGRPCManager.addDownloadTask(r);
    dPrint("[Downloader] addDownloadTask result.id==$id");
    return id;
  }

  static StreamSubscription getDownloadInfoStream(DownloadInfoRequestType type,
      {String? id, void Function(GrpcFileDownloadInfoData value)? onData}) {
    final s = AppGRPCManager.getDownloadInfoStream(
        DownloadInfoRequest(id: id, type: type));
    return s.listen((value) {
      final data = GrpcFileDownloadInfoData.fromJson(
          json.decode(utf8.decode(value.data)));
      // dPrint("[Downloader] info stream :${data.toJson()}");
      onData?.call(data);
    }, onDone: () {}, onError: (error, stackTrace) {}, cancelOnError: true);
  }

  static String getFilePath(String savePath, String fileName) {
    return "$savePath/$fileName";
  }
}
