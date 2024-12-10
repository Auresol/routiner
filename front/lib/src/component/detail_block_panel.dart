import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:routiner/src/model/block.dart';

import 'package:routiner/src/repo/block_repo.dart';
import 'package:provider/provider.dart';

class DetailBlockPanel extends StatefulWidget {

  final Block block;

  DetailBlockPanel({super.key, required this.block});

  @override
  State<DetailBlockPanel> createState() => _DetailBlockPanelState();
}

class _DetailBlockPanelState extends State<DetailBlockPanel> {
  
  bool isDeleting = true;

  @override
  Widget build(BuildContext context) {
    return Builder(
      builder: (context) {
        return Consumer<BlockProvider>(
          builder: (context, provider, child) {
              return Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [
                  Text("Title: ${widget.block.title}"),
                  Text(widget.block.explain ?? 'No explain'),
                  ElevatedButton(
                    child: Text(isDeleting? "Delete" : "Revert"),
                    onPressed: () {
                      if(isDeleting){
                        provider.deleteBlock(widget.block.id!);
                      }else{
                        provider.revertBlock(widget.block.id!);
                      }
                      setState(() {
                        isDeleting = !isDeleting;
                      });
                    }
                  ),
                ],
              );
            }
          
        );
      }
    );
  }
}