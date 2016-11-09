#include "refineimg.h"

#include <opencv2/core.hpp>
#include <opencv2/highgui.hpp>

#include "image_refinement.h"

std::vector<unsigned char> refine(const std::vector<unsigned char>& img) {
  cv::Mat src = cv::imdecode(img, cv::IMREAD_COLOR);
  cv::Mat dst = imgutil::refine(src, NULL);
  std::vector<unsigned char> out;
  cv::imencode(".webp", dst, out);
  return out;
}
