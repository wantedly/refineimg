#ifndef YASHMA_IMAGE_REFINEMENT_H_
#define YASHMA_IMAGE_REFINEMENT_H_

#include <opencv2/opencv.hpp>

namespace imgutil {

cv::Mat refine(const cv::Mat& img, cv::Mat* hist_image = NULL);

}

#endif // YASHMA_IMAGE_REFINEMENT_H_
