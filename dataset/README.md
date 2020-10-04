# RWTH KITTI Semantics Dataset
This dataset was started by Georgios Floros and later extended by other members of the VCI vision group. The annotations have been done using the InteractLabeler 1.2.1 from Cambridge. For questions or comments contact `hermans@vision.rwth-aachen.de` or `osep@vision.rwth-aachen.de`. If you use our dataset please cite the following paper:

```
@inproceedings{Osep16ICRA, 
    title={{Multi-Scale Object Candidates for Generic Object Tracking in Street Scenes}}, 
    author={O\v{s}ep, Aljo\v{s}a and Hermans, Alexander and Engelmann, Francis and Klostermann, Dirk and and Mathias, Markus and Leibe, Bastian}, 
    booktitle={ICRA}, 
    year={2016} 
}
```

## Content
The dataset contains annotations for 203 images from the KITTI visual odometry dataset in the `labels` folder. The labels are not pixel accurate, however we found that they sufficed for all our experiments.

The `splits` folder contains some suggested splits. The train and test split should be obvious. The all 'split' contains all images from the dataset and the all_corrected 'split' excludes a single scene since it overlaps with the tracking training dataset. All sequences come from the odometry set and to the best of our knowledge (apart from sequence 20) there is no overlap with the tracking dataset.

To make it easy the `calibrations` and `images` folders contain the calibration in the default KITTI format and the right rgb images for the ones we labeled here.

When the dataset was first created the sub-sampling of the images from the odometry set went slightly wrong. Meaning the image names are wrong. If you want to compute stereo info, or use temporal context, you can retrieve the name of the original image from the `match_file.txt`.


## Color Coding

| Name       | (r,g,b)         |  7-Class mapping   |
|------------|-----------------|--------------------|
| Car        | (  0,  0,255)   | Object             |
| Road       | (255,  0,  0)   | Road               |
| Mark       | (255,255,  0)   | Road               |
| Building   | (  0,255,  0)   | Building           |
| Sidewalk   | (255,  0,255)   | Road               |
| Tree/Bush  | (  0,255,255)   | Tree/Bush          |
| Pole       | (255,  0,153)   | Sign/Pole          |
| Sign       | (153,  0,255)   | Sign/Pole          |
| Person     | (  0,153,255)   | Object             |
| Wall       | (153,255,  0)   | Building           |
| Sky        | (255,153,  0)   | Sky                |
| Curb       | (  0,255,153)   | Road               |
| Grass/Dirt | (  0,153,153)   | Grass/Dirt         |
| Void       | (  0,  0,  0)   | Void               |


## Extension
In our paper we reported experiments based on the label images from the `labels` folder. After the experiments done for our ICRA16 paper, we annotated further classes and fixed some bugs in the annotation in the `labels_new` folder. We also labeled 5 further images. If you want to compare to our numbers you will have to use the images in the `labels` folder. However, do consider that we only used it as a tool and other papers should be used as proper semantic segmentation baselines. For the extension to the new data these three classes were added, here the 7-Class mapping is not applicable.:

| Name               | (r,g,b)         |
|--------------------|-----------------|
| Side rail          | (153,153,153)   |
| Object             | (  0,  0,153)   |
| Bicycle/Motorbike  | (255,255,153)   |